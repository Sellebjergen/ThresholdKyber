package owcpa_TKyber

import (
	"math/rand"

	"ThresholdKyber.com/m/kyber"
)

// Currently additively secret sharing is hardcoded, would be nice to extract.
func Share(sk kyber.PolyVec, n int) []kyber.PolyVec {
	shares := make([]kyber.PolyVec, n)

	for i := range shares {
		shares[i] = kyber.Kyber512.AllocPolyVec()
	}

	for poly, sk_poly := range sk.Vec {
		poly_shares := SharePolynomial(sk_poly, n)
		for i := 0; i < n; i++ {
			shares[i].Vec[poly] = poly_shares[i]
		}
	}

	return shares
}

func Rec(d_is [][]*kyber.Poly, r int) kyber.PolyVec {
	recombined := kyber.Kyber512.AllocPolyVec()
	for _, d_i := range d_is {
		recombined.Vec = append(recombined.Vec, RecPolynomial(d_i))
	}
	return recombined
}

func RecPolynomial(d_is []*kyber.Poly) *kyber.Poly {
	var out kyber.Poly

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
		shares[i] = SampleUnifPolynomial(3329) // TODO: Kyber params
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
