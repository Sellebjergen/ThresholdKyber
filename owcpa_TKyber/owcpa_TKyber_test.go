package owcpa_TKyber

import (
	"reflect"
	"testing"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
)

// ================= Setup tests =================

func TestSetupWorksInCaseNis3(t *testing.T) {
	params := NewParameterSet("TKyber-Test")
	pk, sk_shares := Setup(params, 3, 3)

	// total of first share
	var sk1 kyberk2so.Poly
	sk1 = kyberk2so.PolyAdd(sk1, sk_shares[0][0])
	sk1 = kyberk2so.PolyAdd(sk1, sk_shares[1][0])
	sk1 = kyberk2so.PolyAdd(sk1, sk_shares[2][0])

	// total of second share
	var sk2 kyberk2so.Poly
	sk2 = kyberk2so.PolyAdd(sk2, sk_shares[0][1])
	sk2 = kyberk2so.PolyAdd(sk2, sk_shares[1][1])
	sk2 = kyberk2so.PolyAdd(sk2, sk_shares[2][1])

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

// ================= Integration tests =================

func TestAdvancedCase(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	params := NewParameterSet("TKyber-Test")
	n := 4
	t_param := 4
	pk, sk_shares := Setup(params, n, t_param)

	ct := Enc(params, msg, pk)

	// Decrypt
	d_is := make([]kyberk2so.Poly, n)
	for i := 0; i < n; i++ {
		d_is[i] = PartDec(params, sk_shares[i], ct, i)
	}

	combined := Combine(params, ct, d_is)
	output_msg := kyberk2so.PolyToMsg(combined)

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Error")
	}
}

func TestSimpleCase(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	params := NewParameterSet("TKyber-Test")
	pk, sk_shares := Setup(params, 3, 3)

	coins := make([]byte, 32)
	// rand.Read(coins)

	ct, _ := kyberk2so.IndcpaEncrypt(msg, pk, coins, kyberk2so.ParamsK)

	// Decrypt
	d_is := make([]kyberk2so.Poly, 3)
	d_is[0] = PartDec(params, sk_shares[0], ct, 0)
	d_is[1] = PartDec(params, sk_shares[1], ct, 1)
	d_is[2] = PartDec(params, sk_shares[2], ct, 2)

	combined := Combine(params, ct, d_is)
	output_msg := kyberk2so.PolyToMsg(combined)

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Error")
	}
}

func TestSetupUsing1PlayerGivesBackSecretKey(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	params := NewParameterSet("TKyber-Test")
	pk, skShares := Setup(params, 1, 1)
	polyVec := kyberk2so.PolyvecNew(kyberk2so.ParamsK)
	polyVec[0] = skShares[0][0]
	polyVec[1] = skShares[0][1]

	coins := make([]byte, 32)
	ct, _ := kyberk2so.IndcpaEncrypt(msg, pk, coins, kyberk2so.ParamsK)

	sk_packed := kyberk2so.IndcpaPackPrivateKey(polyVec, kyberk2so.ParamsK)
	out := kyberk2so.IndcpaDecrypt(ct, sk_packed, kyberk2so.ParamsK)

	if !reflect.DeepEqual(msg, out) {
		t.Errorf("Error")
	}
}

func TestFullWithN1(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	params := NewParameterSet("TKyber-Test")
	pk, skShares := Setup(params, 1, 1)

	coins := make([]byte, 32)

	ct, _ := kyberk2so.IndcpaEncrypt(msg, pk, coins, kyberk2so.ParamsK)

	d_is := make([]kyberk2so.Poly, 1)
	d_is[0] = PartDec(params, skShares[0], ct, 0)

	combined := Combine(params, ct, d_is)
	output_msg := kyberk2so.PolyToMsg(combined)

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Error")
	}
}
