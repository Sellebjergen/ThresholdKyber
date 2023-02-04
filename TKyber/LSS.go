package TKyber

import (
	"math/rand"
)

// Represents additively secret sharing
func (rq *quotRing) Share(sk []*Polynomial, n int) [][]*Polynomial {
	r := len(sk)
	shares := make([][]*Polynomial, n)
	for i := range shares {
		shares[i] = make([]*Polynomial, r)
	}
	for poly, sk_poly := range sk {
		poly_shares := rq.SharePolynomial(sk_poly, n)
		for i := 0; i < n; i++ {
			shares[i][poly] = poly_shares[i]
		}
	}

	return shares
}

func (rq *quotRing) Rec(d_is [][]*Polynomial, r int) []*Polynomial {
	recombined := make([]*Polynomial, 0)
	for _, d_i := range d_is {
		recombined = append(recombined, rq.RecPolynomial(d_i))
	}
	return recombined
}

func (rq *quotRing) RecPolynomial(d_is []*Polynomial) *Polynomial {
	out := d_is[0]

	for i := 1; i < len(d_is); i++ {
		out = rq.add(out, d_is[i])
	}

	return out
}

func (rq *quotRing) SharePolynomial(toShare *Polynomial, n int) []*Polynomial {
	shares := make([]*Polynomial, n)

	for i := 0; i <= n-2; i++ {
		shares[i] = SampleUnifPolynomial(3329, 256) // TODO: Kyber params
	}

	shares[n-1] = toShare.Copy()
	for i := 0; i <= n-2; i++ {
		shares[n-1] = rq.sub(shares[n-1], shares[i])
	}

	return shares
}

func SampleUnifPolynomial(q int, deg int) *Polynomial {
	coeffs := make([]int, deg+1)
	for i := 0; i < deg+1; i++ {
		coeffs[i] = rand.Intn(q) // TODO: Kyber params
	}
	return &Polynomial{Coeffs: coeffs}
}
