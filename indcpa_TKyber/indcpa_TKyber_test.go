package indcpa_TKyber

import (
	"crypto/rand"
	"fmt"
	"reflect"
	"testing"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
	owcpa "ThresholdKyber.com/m/owcpa_TKyber"
)

// ================= Integration tests =================

func TestSimpleCase(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	params := owcpa.NewParameterSet("TKyber-Test")
	n := 1
	t_param := 1
	delta := 2
	pk, sk_shares := Setup(params, n, t_param)

	ct := Enc(params, msg, pk, delta)

	// Decrypt
	d_1 := PartDec(params, sk_shares[0], ct, 0, delta)

	d_is := [][][]kyberk2so.Poly{d_1}

	combined := Combine(params, ct, d_is)

	output_msg := kyberk2so.PolyToMsg(combined)

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Error")
	}
}

func TestAdvancedCase(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	params := owcpa.NewParameterSet("TKyber-Test")
	n := 3
	t_param := 3
	delta := 10
	pk, sk_shares := Setup(params, n, t_param)

	ct := Enc(params, msg, pk, delta)

	// Decrypt
	d_1 := PartDec(params, sk_shares[0], ct, 0, delta)
	d_2 := PartDec(params, sk_shares[1], ct, 1, delta)
	d_3 := PartDec(params, sk_shares[2], ct, 2, delta)

	d_is := [][][]kyberk2so.Poly{d_1, d_2, d_3}

	combined := Combine(params, ct, d_is)

	output_msg := kyberk2so.PolyToMsg(combined)

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Error")
	}
}

func TestLowDeltaCase(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	params := owcpa.NewParameterSet("TKyber-Test")
	n := 3
	t_param := 3
	delta := 1
	pk, sk_shares := Setup(params, n, t_param)

	ct := Enc(params, msg, pk, delta)

	// Decrypt
	d_1 := PartDec(params, sk_shares[0], ct, 0, delta)
	d_2 := PartDec(params, sk_shares[1], ct, 1, delta)
	d_3 := PartDec(params, sk_shares[2], ct, 2, delta)

	d_is := [][][]kyberk2so.Poly{d_1, d_2, d_3}

	combined := Combine(params, ct, d_is)

	output_msg := kyberk2so.PolyToMsg(combined)

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Error")
	}
}

func BenchmarkSetupTKyber(b *testing.B) {
	cases := []struct {
		TKyberVariant string
		n             int
		t             int
		delta         int
	}{
		{TKyberVariant: "TKyber-Test", n: 1, t: 1, delta: 10},
		{TKyberVariant: "TKyber-Test", n: 2, t: 2, delta: 10},
		{TKyberVariant: "TKyber-Test", n: 3, t: 3, delta: 10},
		{TKyberVariant: "TKyber-Test", n: 1, t: 1, delta: 1000},
	}

	for _, bCase := range cases {
		b.Run(fmt.Sprintf("Setup %s", bCase.TKyberVariant), func(b *testing.B) { benchmarkSetup(b, bCase.TKyberVariant, bCase.n, bCase.t) })
		b.Run(fmt.Sprintf("Enc %s", bCase.TKyberVariant), func(b *testing.B) { benchmarkEnc(b, bCase.TKyberVariant, bCase.n, bCase.t, bCase.delta) })
		b.Run(fmt.Sprintf("PartDec %s", bCase.TKyberVariant), func(b *testing.B) { benchmarkPartDec(b, bCase.TKyberVariant, bCase.n, bCase.t, bCase.delta) })
		b.Run(fmt.Sprintf("Combine %s", bCase.TKyberVariant), func(b *testing.B) { benchmarkCombine(b, bCase.TKyberVariant, bCase.n, bCase.t, bCase.delta) })
	}
}

func benchmarkSetup(b *testing.B, paramSet string, n int, t int) {
	params := owcpa.NewParameterSet(paramSet)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Setup(params, n, t)
	}
}

func benchmarkEnc(b *testing.B, paramSet string, n int, t int, delta int) {
	randMsg := make([]byte, 32)
	rand.Read(randMsg)
	params := owcpa.NewParameterSet(paramSet)
	pk, _ := Setup(params, n, t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Enc(params, randMsg, pk, delta)
	}
}

func benchmarkPartDec(b *testing.B, paramSet string, n int, t int, delta int) {
	randMsg := make([]byte, 32)
	rand.Read(randMsg)
	params := owcpa.NewParameterSet(paramSet)
	pk, sk_shares := Setup(params, n, t)
	ct := Enc(params, randMsg, pk, delta)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PartDec(params, sk_shares[0], ct, 0, delta) // Er det fint med party 0 her?
	}
}

func benchmarkCombine(b *testing.B, paramSet string, n int, t int, delta int) {
	randMsg := make([]byte, 32)
	rand.Read(randMsg)
	params := owcpa.NewParameterSet(paramSet)
	pk, sk_shares := Setup(params, n, t)
	ct := Enc(params, randMsg, pk, delta)

	d_is := make([][][]kyberk2so.Poly, 0)
	for i := 0; i < t; i++ {
		d_is = append(d_is, PartDec(params, sk_shares[i], ct, i, delta))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Combine(params, ct, d_is)
	}
}
