package TKyber

import (
	"crypto/rand"
	"math"

	"ThresholdKyber.com/m/kyber"
)

func Setup(params kyber.ParameterSet, n int, t int) (*kyber.IndcpaPublicKey, []kyber.PolyVec) {
	// Run setup to get Kyber KeyPair
	pk, sk, _ := params.IndcpaKeyPair(rand.Reader)
	polyVec_sk := params.AllocPolyVec()
	kyber.UnpackSecretKey(&polyVec_sk, sk.Packed)

	// Perform secret sharing
	sk_shares := Share(polyVec_sk, n)

	return pk, sk_shares
}

func PartDec(params kyber.ParameterSet, sk_i kyber.PolyVec, ct []byte, party int) *kyber.Poly {
	var v, d_i, zero kyber.Poly
	// Sample noise
	//e_i := samplePolyGaussian(3329, 255, 0) // TODO: Fix params

	// Convert bytes from ct to list of polynomials (internal type)
	bp := params.AllocPolyVec()
	kyber.UnpackCiphertext(&bp, &v, ct) // This will be NTT form.

	// Inner prod
	bp.Ntt()
	d_i.PointwiseAcc(&sk_i, &bp)
	d_i.Invntt()

	if party == 0 {
		d_i.Sub(&v, &d_i)
	} else {
		d_i.Sub(&zero, &d_i)
	}

	// Add noise
	//d_i = rq.add(d_i, e_i)

	return &d_i
}

func (rq *quotRing) Combine(ct []byte, d_is ...*kyber.Poly) *kyber.Poly {
	p := 2
	y := RecPolynomial(d_is)
	unrounded := make([]float64, len(y.Coeffs))

	for i := 0; i < len(unrounded); i++ {
		unrounded[i] = (float64(p) / float64(rq.q)) * float64(y.Coeffs[i])
	}

	res := make([]uint16, len(y.Coeffs))
	for i := 0; i < len(unrounded); i++ {
		res[i] = uint16(math.Round(unrounded[i]))
	}
	out := new(kyber.Poly)
	copy(out.Coeffs[:], res)

	for i := 0; i < len(out.Coeffs); i++ {
		out.Coeffs[i] = kyber.Freeze(out.Coeffs[i])
	}

	return out
}
