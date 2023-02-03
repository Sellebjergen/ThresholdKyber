package TKyber

import (
	"crypto/rand"

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

func PartDec() {

}

func (r *quotRing) Combine(ct []byte, d_is []*Polynomial) *kyber.Poly {
	r.RecPolynomial(d_is) // Can't use return value at the moment
	return nil
}
