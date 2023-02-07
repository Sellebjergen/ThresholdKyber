package indcpa_TKyber

import (
	"ThresholdKyber.com/m/kyber"
	owcpa "ThresholdKyber.com/m/owcpa_TKyber"
)

func Setup(params *owcpa.OwcpaParams, n int, t int) (*kyber.IndcpaPublicKey, []kyber.PolyVec) {
	return owcpa.Setup(params, n, t)
}

func Enc(params *owcpa.OwcpaParams, msg []byte, pk *kyber.IndcpaPublicKey, delta int) {
	var mp kyber.Poly
	mp.ToMsg(msg)
	x := make([]*kyber.Poly, delta)
	for i := 0; i < delta; i++ {
		x[i] = owcpa.SampleUnifPolynomial(2)
	}
	c := make([]*kyber.Poly, delta+1)
	c[0].Add(&mp, F(x))

}

func PartDec() {

}

func Combine() {

}

func F(x []*kyber.Poly) *kyber.Poly {
	return nil
}

func G(x []*kyber.Poly) []byte {
	return nil
}
