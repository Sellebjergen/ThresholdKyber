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

// ================= Testing that Share on polynomial ring ====
func TestShare(t *testing.T) {
	sk := make([]*Polynomial, 2)
	sk[0] = &Polynomial{Coeffs: []int{42, 73}}
	sk[1] = &Polynomial{Coeffs: []int{0, 27}}

	rq := new(quotRing).initKyberRing()
	shares := rq.Share(sk, 3)

	// combine the first polynomial
	sk1 := &Polynomial{[]int{0}}
	sk1 = rq.add(sk1, shares[0][0])
	sk1 = rq.add(sk1, shares[1][0])
	sk1 = rq.add(sk1, shares[2][0])

	// combine the second polynomial
	sk2 := &Polynomial{[]int{0}}
	sk2 = rq.add(sk2, shares[0][1])
	sk2 = rq.add(sk2, shares[1][1])
	sk2 = rq.add(sk2, shares[2][1])

	if !reflect.DeepEqual(sk[0], sk1) {
		t.Errorf("The first shared polynomial is not equal expected!")
	}

	if !reflect.DeepEqual(sk[1], sk2) {
		t.Errorf("The second shared polynomial is not equal expected!")
	}
}
