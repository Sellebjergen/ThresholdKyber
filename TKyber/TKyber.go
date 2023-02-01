package TKyber

import (
	"crypto/rand"

	"ThresholdKyber.com/m/kyber"
)

func Setup(params kyber.ParameterSet, n int, t int) (*kyber.IndcpaPublicKey, []*share) {
	pk, sk, _ := params.IndcpaKeyPair(rand.Reader)
	sk_shares := Share(params, sk)
	return pk, sk_shares
}

func PartDec() {

}

func (r *polyRing) Combine(ct []byte, d_is []*share) *kyber.Poly {
	r.Rec(d_is) // Can't use return value at the moment
	return nil
}
