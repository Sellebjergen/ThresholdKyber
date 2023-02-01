package TKyber

import (
	"crypto/rand"

	"ThresholdKyber.com/m/kyber"
)

func Setup(params kyber.ParameterSet, n int, t int) (*kyber.IndcpaPublicKey, []*SKshare) {
	pk, sk, _ := params.IndcpaKeyPair(rand.Reader)
	sk_shares := Share(params, sk)
	return pk, sk_shares
}

func PartDec() {

}

func Combine() {

}
