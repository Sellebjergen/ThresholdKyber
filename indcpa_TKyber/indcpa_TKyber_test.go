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

func TestTKyberConsistency(t *testing.T) {
	cases := []struct {
		TKyberVariant string
		n             int
		t             int
		delta         int
	}{

		{TKyberVariant: "TKyber-Test", n: 1, t: 0, delta: 2},
		{TKyberVariant: "TKyber-Test", n: 2, t: 1, delta: 2},
		{TKyberVariant: "TKyber-Test", n: 3, t: 2, delta: 2},
		{TKyberVariant: "TKyber-Test", n: 4, t: 3, delta: 2},

		{TKyberVariant: "TKyber-Test", n: 3, t: 2, delta: 1},
		{TKyberVariant: "TKyber-Test", n: 3, t: 2, delta: 10},

		{TKyberVariant: "TKyber-Test-Replicated", n: 3, t: 1, delta: 2},
		{TKyberVariant: "TKyber-Test-Replicated", n: 3, t: 2, delta: 2},

		{TKyberVariant: "TKyber-Test-Naive", n: 3, t: 1, delta: 2},
		{TKyberVariant: "TKyber-Test-Naive", n: 3, t: 2, delta: 2},
	}

	for _, tCase := range cases {
		t.Run(fmt.Sprintf("%s with n = %d, t = %d, delta = %d", tCase.TKyberVariant, tCase.n, tCase.t, tCase.delta),
			func(t *testing.T) { testConsistencyCheck(t, tCase.TKyberVariant, tCase.n, tCase.t, tCase.delta) })
	}
}

func testConsistencyCheck(t *testing.T, TKyberVariant string, n, t_param, delta int) {
	msg := make([]byte, 32)
	rand.Read(msg)
	params := owcpa.NewParameterSet(TKyberVariant)
	pk, sk_shares := Setup(params, n, t_param)

	ct := Enc(params, msg, pk, delta)

	// Decrypt
	d_is := make([][][]kyberk2so.Poly, n)
	for i := 0; i < t_param+1; i++ {
		d_is[i] = PartDec(params, sk_shares[i], ct, i, delta)
	}
	for i := t_param + 1; i < n; i++ {
		d_is[i] = make([][]kyberk2so.Poly, len(sk_shares[i]))
	}

	combined := Combine(params, ct, d_is, n, t_param)
	output_msg := kyberk2so.PolyToMsg(combined)

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Consistency test failed")
	}
}

// ================= Benchmarking =================

func BenchmarkTKyber(b *testing.B) {
	cases := []struct {
		TKyberVariant string
		n             int
		t             int
		delta         int
	}{
		{TKyberVariant: "TKyber-Test", n: 1, t: 1, delta: 1},
		{TKyberVariant: "TKyber-Test", n: 2, t: 2, delta: 1},
		{TKyberVariant: "TKyber-Test", n: 3, t: 3, delta: 1},
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
		Combine(params, ct, d_is, n, t)
	}
}
