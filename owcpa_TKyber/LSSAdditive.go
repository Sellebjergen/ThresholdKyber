package owcpa_TKyber

import (
	"math/rand"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
)

type LSSAdditive struct{}

// Currently additively secret sharing is hardcoded, would be nice to extract.
func (s *LSSAdditive) Share(sk kyberk2so.PolyVec, n int) []kyberk2so.PolyVec {
	shares := make([]kyberk2so.PolyVec, n)
	paramsK := 2

	for i := range shares {
		shares[i] = kyberk2so.PolyvecNew(paramsK)
	}

	for poly, sk_poly := range sk {
		poly_shares := SharePolynomial(sk_poly, n)
		for i := 0; i < n; i++ {
			shares[i][poly] = poly_shares[i]
		}
	}

	return shares
}

func (s *LSSAdditive) Rec(d_is []kyberk2so.Poly) kyberk2so.Poly {
	var out kyberk2so.Poly

	for i := 0; i < len(d_is); i++ {
		out = kyberk2so.PolyAdd(out, d_is[i])
	}
	out = kyberk2so.PolyReduce(out)

	return out
}

func SharePolynomial(toShare kyberk2so.Poly, n int) []kyberk2so.Poly {
	shares := make([]kyberk2so.Poly, n)

	for i := 0; i <= n-2; i++ {
		shares[i] = SampleUnifPolynomial(7681) // TODO: Kyber params
	}

	shares[n-1] = Copy(toShare)
	for i := 0; i <= n-2; i++ {
		shares[n-1] = kyberk2so.PolySub(shares[n-1], shares[i])
	}

	return shares
}

func SampleUnifPolynomial(q int) kyberk2so.Poly {
	var out_coeff [kyberk2so.ParamsPolyBytes]int16
	for i := 0; i < 256; i++ {
		out_coeff[i] = int16(rand.Intn(q)) // TODO: Kyber params
	}
	return kyberk2so.Poly(out_coeff)
}

func Copy(toCopy kyberk2so.Poly) kyberk2so.Poly {
	var out kyberk2so.Poly
	coeff := toCopy
	copy(out[:], coeff[:])

	return out
}
