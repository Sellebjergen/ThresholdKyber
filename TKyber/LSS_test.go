package TKyber

import (
	"reflect"
	"testing"
)

// ================= Share tests =================
func TestSharePolynomial(t *testing.T) {
	toShare := &Polynomial{Coeffs: []int{1, 2, 3, 4, 5}}

	rq := new(quotRing).initKyberRing()
	shares := rq.SharePolynomial(toShare, 10)

	recombined := &Polynomial{Coeffs: []int{0}}

	for _, share := range shares {
		recombined = rq.add(recombined, share)
	}

	if !reflect.DeepEqual(recombined, toShare) {
		t.Errorf("Recombined is not equal the original shared polynomial!")
	}
}

// ================= Share then Rec polynomial test =================

func TestShareThenRecPolynomial(t *testing.T) {
	toShare := &Polynomial{Coeffs: []int{1, 2, 3, 4, 5}}

	rq := new(quotRing).initKyberRing()
	shares := rq.SharePolynomial(toShare, 10)

	recombined := rq.RecPolynomial(shares)

	if !reflect.DeepEqual(recombined, toShare) {
		t.Errorf("Recombined is not equal the original shared polynomial!")
	}
}
