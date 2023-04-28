package owcpa_TKyber

import (
	"crypto/rand"
	"fmt"
	"reflect"
	"testing"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
)

// ================= Integration tests =================

func TestTKyberConsistency(t *testing.T) {
	cases := []struct {
		TKyberVariant string
		n             int
		t             int
	}{
		{TKyberVariant: "TKyber-Test", n: 1, t: 0},
		{TKyberVariant: "TKyber-Test", n: 2, t: 1},
		{TKyberVariant: "TKyber-Test", n: 3, t: 2},
		{TKyberVariant: "TKyber-Test", n: 4, t: 3},

		{TKyberVariant: "TKyber-Test-Replicated", n: 3, t: 1},
		{TKyberVariant: "TKyber-Test-Replicated", n: 3, t: 2},

		{TKyberVariant: "TKyber-Test-Naive", n: 3, t: 1},
		{TKyberVariant: "TKyber-Test-Naive", n: 3, t: 2},
	}

	for _, tCase := range cases {
		t.Run(fmt.Sprintf("%s with n = %d, t = %d", tCase.TKyberVariant, tCase.n, tCase.t),
			func(t *testing.T) { testConsistencyCheck(t, tCase.TKyberVariant, tCase.n, tCase.t) })
	}
}

func testConsistencyCheck(t *testing.T, TKyberVariant string, n, t_param int) {
	msg := make([]byte, 32)
	rand.Read(msg)
	params := NewParameterSet(TKyberVariant)
	pk, sk_shares := Setup(params, n, t_param)

	ct := Enc(params, msg, pk)

	// Decrypt
	d_is := make([][]kyberk2so.Poly, n)
	for i := 0; i < t_param+1; i++ {
		d_is[i] = PartDec(params, sk_shares[i], ct, i)
	}

	combined := Combine(params, ct, d_is, n, t_param)
	output_msg := kyberk2so.PolyToMsg(combined)

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Consistency test failed")
	}
}

func TestFullDeterministic(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	params := NewParameterSet("TKyber-Test")
	pk, sk_shares := Setup(params, 3, 2)

	coins := make([]byte, 32)
	// rand.Read(coins)

	ct, _ := kyberk2so.IndcpaEncrypt(msg, pk, coins, kyberk2so.ParamsK)

	// Decrypt
	d_is := make([][]kyberk2so.Poly, 3)
	d_is[0] = PartDec(params, sk_shares[0], ct, 0)
	d_is[1] = PartDec(params, sk_shares[1], ct, 1)
	d_is[2] = PartDec(params, sk_shares[2], ct, 2)

	combined := Combine(params, ct, d_is, 3, 2)
	output_msg := kyberk2so.PolyToMsg(combined)

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Error")
	}
}

// This represents an old bug found that used to be in the IND-CPA TKyber code.
func TestSimINDCPATransform(t *testing.T) {
	params := NewParameterSet("TKyber-Test")
	pk, skShares := Setup(params, 1, 0)
	m := SampleUnifPolynomial(2)
	upscaled := Upscale(m, 2, params.Q)
	m_bytes := kyberk2so.PolyToMsg(upscaled)

	coins := make([]byte, 32)
	ct, _ := kyberk2so.IndcpaEncrypt(m_bytes, pk, coins, kyberk2so.ParamsK)
	part := PartDec(params, skShares[0], ct, 0)

	res := Combine(params, ct, [][]kyberk2so.Poly{part}, 1, 0)
	downscaled := Downscale(res, 2, params.Q)

	if !reflect.DeepEqual(downscaled, m) {
		t.Errorf("Error: Polynomials not matching")
	}
}

// ================= Setup tests =================

func TestSetupWorksInCaseNis3(t *testing.T) {
	params := NewParameterSet("TKyber-Test")
	pk, sk_shares := Setup(params, 3, 3)

	// total of first share
	var sk1 kyberk2so.Poly
	sk1 = kyberk2so.PolyAdd(sk1, sk_shares[0][0][0])
	sk1 = kyberk2so.PolyAdd(sk1, sk_shares[1][0][0])
	sk1 = kyberk2so.PolyAdd(sk1, sk_shares[2][0][0])

	// total of second share
	var sk2 kyberk2so.Poly
	sk2 = kyberk2so.PolyAdd(sk2, sk_shares[0][0][1])
	sk2 = kyberk2so.PolyAdd(sk2, sk_shares[1][0][1])
	sk2 = kyberk2so.PolyAdd(sk2, sk_shares[2][0][1])

	// assemble secret key
	sk := kyberk2so.PolyVec{sk1, sk2}

	sk_packed := kyberk2so.IndcpaPackPrivateKey(sk, kyberk2so.ParamsK)

	// calling the encrypt decrypt functionality, to check that the key works as expected.
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	coins := make([]byte, 32)

	ct, _ := kyberk2so.IndcpaEncrypt(msg, pk, coins, kyberk2so.ParamsK)
	out := kyberk2so.IndcpaDecrypt(ct, sk_packed, kyberk2so.ParamsK)

	if !reflect.DeepEqual(msg, out) {
		t.Errorf("Decryption failed!")
	}
}

func TestSetupUsing1PlayerGivesBackSecretKey(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	params := NewParameterSet("TKyber-Test")
	pk, skShares := Setup(params, 1, 1)
	polyVec := []kyberk2so.PolyVec{kyberk2so.PolyvecNew(kyberk2so.ParamsK)}
	polyVec[0][0] = skShares[0][0][0]
	polyVec[0][1] = skShares[0][0][1]

	coins := make([]byte, 32)
	ct, _ := kyberk2so.IndcpaEncrypt(msg, pk, coins, kyberk2so.ParamsK)

	sk_packed := kyberk2so.IndcpaPackPrivateKey(polyVec[0], kyberk2so.ParamsK)
	out := kyberk2so.IndcpaDecrypt(ct, sk_packed, kyberk2so.ParamsK)

	if !reflect.DeepEqual(msg, out) {
		t.Errorf("Error")
	}
}

// ================= Replicated LSS tests =================

func TestWithReplicatedLSSNoCombine(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	params := NewParameterSet("TKyber-Test-Replicated")
	n := 3
	t_param := 1
	pk, sk_shares := Setup(params, n, t_param)

	ct := Enc(params, msg, pk)

	// Decrypt
	d_1 := PartDec(params, sk_shares[0], ct, 0)
	d_2 := PartDec(params, sk_shares[1], ct, 1)
	d_3 := PartDec(params, sk_shares[2], ct, 2)

	d_is := [][]kyberk2so.Poly{d_1, d_2, d_3}

	var p kyberk2so.Poly

	p = kyberk2so.PolyAdd(p, d_is[0][1])
	p = kyberk2so.PolyAdd(p, d_is[0][2])
	p = kyberk2so.PolyAdd(p, d_is[1][0])

	p = kyberk2so.PolyReduce(p)

	output_msg := kyberk2so.PolyToMsg(p)

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Error")
	}
}

// ================= Binomial noise tests =================
func TestWithBinomialNoise(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	params := NewParameterSet("TKyber-Test")
	params.D_flood_dist = &BinomialNoiseDist{Eta: 2}
	pk, sk_shares := Setup(params, 2, 2)

	for i := 0; i < 100; i++ {
		ct := Enc(params, msg, pk)

		// Decrypt
		d_is := make([][]kyberk2so.Poly, 2)
		d_is[0] = PartDec(params, sk_shares[0], ct, 0)
		d_is[1] = PartDec(params, sk_shares[1], ct, 1)

		combined := Combine(params, ct, d_is, 2, 2)
		output_msg := kyberk2so.PolyToMsg(combined)

		if !reflect.DeepEqual(msg, output_msg) {
			t.Errorf("Error")
		}
	}

}
