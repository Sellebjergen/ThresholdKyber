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

func (s *LSSAdditive) Rec(d_is []*kyber.Poly) *kyber.Poly {
	var out kyberk2so.Poly

	for i := 0; i < len(d_is); i++ {
		out.Add(&out, d_is[i])
	}
	for i := 0; i < len(out.Coeffs); i++ {
		out.Coeffs[i] = kyber.Freeze(out.Coeffs[i])
	}

	return &out
}

func SharePolynomial(toShare *kyber.Poly, n int) []*kyber.Poly {
	shares := make([]*kyber.Poly, n)

	for i := 0; i <= n-2; i++ {
		shares[i] = SampleUnifPolynomial(7681) // TODO: Kyber params
	}

	shares[n-1] = Copy(toShare)
	for i := 0; i <= n-2; i++ {
		shares[n-1].Sub(shares[n-1], shares[i])
	}

	return shares
}

func SampleUnifPolynomial(q int) *kyber.Poly {
	var out_coeff [256]uint16
	for i := 0; i < 256; i++ {
		out_coeff[i] = uint16(rand.Intn(q)) // TODO: Kyber params
	}
	return &kyber.Poly{Coeffs: out_coeff}
}

func Copy(toCopy *kyber.Poly) *kyber.Poly {
	var out kyber.Poly
	coeff := toCopy.Coeffs
	copy(out.Coeffs[:], coeff[:])

	return &out
}
