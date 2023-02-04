package TKyber

import (
	"crypto/rand"
	"fmt"
	"math"

	"ThresholdKyber.com/m/kyber"
)

func Setup(params kyber.ParameterSet, n int, t int) (*kyber.IndcpaPublicKey, [][]*Polynomial) {
	// Run setup to get Kyber KeyPair
	pk, sk, _ := params.IndcpaKeyPair(rand.Reader)
	polyVec_sk := params.AllocPolyVec()
	kyber.UnpackSecretKey(&polyVec_sk, sk.Packed)

	// Covert from Kyber poly to internal Polynomial type
	sk_internal := make([]*Polynomial, len(polyVec_sk.Vec))
	for i, poly := range polyVec_sk.Vec {
		sk_internal[i] = fromKyberPoly(poly)
	}

	// Perform secret sharing
	rq := new(quotRing).initKyberRing()
	sk_shares := rq.Share(sk_internal, n)

	return pk, sk_shares
}

func (rq *quotRing) PartDec(params kyber.ParameterSet, sk_i []*Polynomial, ct []byte, party int) *Polynomial {

	// Sample noise
	e_i := samplePolyGaussian(3329, 255, 0) // TODO: Fix params

	// Convert bytes from ct to list of polynomials (internal type)
	b := params.AllocPolyVec()
	v := new(kyber.Poly)
	v_internal := fromKyberPoly(v)
	kyber.UnpackCiphertext(&b, v, ct)
	ct_as_internal := make([]*Polynomial, len(b.Vec))
	for i := 0; i < len(b.Vec); i++ {
		ct_as_internal[i] = fromKyberPoly(b.Vec[i])
	}

	// Inner prod
	d_i := &Polynomial{Coeffs: []int{0}}
	fmt.Println(len(ct_as_internal))
	fmt.Println(len(sk_i))
	for poly := 0; poly < len(b.Vec); poly++ {
		inner_prod_part := rq.mult(ct_as_internal[poly], sk_i[poly])
		d_i = rq.add(d_i, inner_prod_part)
	}
	if party == 0 {
		d_i = rq.sub(v_internal, d_i)
	} else {
		d_i = rq.neg(d_i)
	}

	// Add noise
	d_i = rq.add(d_i, e_i)

	// Return d_i
	return d_i
}

func (rq *quotRing) Combine(ct []byte, d_is ...*Polynomial) *kyber.Poly {
	p := 2                      // WAT?
	y := rq.RecPolynomial(d_is) // Can't use return value at the moment
	unrounded := make([]float64, len(y.Coeffs))
	for i := 0; i < len(unrounded); i++ {
		unrounded[i] = float64(p) / float64(rq.q) * float64(y.Coeffs[i])
	}

	res := make([]int, len(y.Coeffs))
	for i := 0; i < len(unrounded); i++ {
		res[i] = int(math.Round(unrounded[i]))
	}
	internal_poly := &Polynomial{Coeffs: res}

	return internal_poly.toKyberPoly()
}
