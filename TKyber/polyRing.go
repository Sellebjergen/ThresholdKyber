package TKyber

import "ThresholdKyber.com/m/kyber"

type polyRing struct {
	q   int
	mod *kyber.Poly
}

func (*polyRing) init() *polyRing {
	rq := &polyRing{
		q:   3329,
		mod: getModuloPoly(),
	}

	return rq
}

func (*polyRing) add() {

}

func (*polyRing) sub() {

}

func (*polyRing) mult() {

}

func (*polyRing) polynomialLongDivision(pol kyber.Poly) {
	// TODO
}

func (*polyRing) syntheticLongDivison(pol kyber.Poly) {

}

func getModuloPoly() *kyber.Poly {
	res := [256]uint16{}
	res[0] = 1
	res[255] = 1

	return &kyber.Poly{Coeffs: res}
}
