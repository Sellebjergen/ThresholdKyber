package owcpa_TKyber

import (
	"fmt"
	"reflect"
	"testing"

	"ThresholdKyber.com/m/kyber"
)

// ================= Share tests =================
func TestSharePolynomial(t *testing.T) {
	var toShare kyber.Poly
	toShare.Coeffs[0] = 1
	toShare.Coeffs[1] = 2
	toShare.Coeffs[2] = 3
	toShare.Coeffs[3] = 4
	toShare.Coeffs[4] = 5

	shares := SharePolynomial(&toShare, 10)

	var recombined kyber.Poly

	for _, share := range shares {
		recombined.Add(&recombined, share)
	}

	for i := 0; i < len(recombined.Coeffs); i++ {
		recombined.Coeffs[i] = kyber.Freeze(recombined.Coeffs[i])
	}

	if !reflect.DeepEqual(recombined.Coeffs, toShare.Coeffs) {
		t.Errorf("Recombined is not equal the original shared polynomial!")
	}
}

// ================= Share then Rec polynomial test =================
func TestShareThenRecPolynomial(t *testing.T) {
	lss := &LSSAdditive{}
	var toShare kyber.Poly
	toShare.Coeffs[0] = 1
	toShare.Coeffs[1] = 2
	toShare.Coeffs[2] = 3
	toShare.Coeffs[3] = 4
	toShare.Coeffs[4] = 5
	fmt.Println(toShare)

	shares := SharePolynomial(&toShare, 10)

	recombined := lss.Rec(shares)

	if !reflect.DeepEqual(recombined.Coeffs, toShare.Coeffs) {
		t.Errorf("Recombined is not equal the original shared polynomial!")
	}
}

// ================= Testing that Share on polynomial ring ====
func TestShare(t *testing.T) {
	sk := kyber.Kyber512.AllocPolyVec()
	lss := &LSSAdditive{}
	sk.Vec[0] = new(kyber.Poly)
	sk.Vec[0].Coeffs[0] = 42
	sk.Vec[0].Coeffs[1] = 73
	sk.Vec[1] = new(kyber.Poly)
	sk.Vec[1].Coeffs[0] = 0
	sk.Vec[1].Coeffs[1] = 27
	shares := lss.Share(sk, 3)

	// combine the first polynomial
	var sk1 kyber.Poly
	sk1.Add(&sk1, shares[0].Vec[0])
	sk1.Add(&sk1, shares[1].Vec[0])
	sk1.Add(&sk1, shares[2].Vec[0])

	// combine the second polynomial
	var sk2 kyber.Poly
	sk2.Add(&sk2, shares[0].Vec[1])
	sk2.Add(&sk2, shares[1].Vec[1])
	sk2.Add(&sk2, shares[2].Vec[1])

	for i := 0; i < len(sk1.Coeffs); i++ {
		sk1.Coeffs[i] = kyber.Freeze(sk1.Coeffs[i])
	}

	for i := 0; i < len(sk2.Coeffs); i++ {
		sk2.Coeffs[i] = kyber.Freeze(sk2.Coeffs[i])
	}

	if !reflect.DeepEqual(sk.Vec[0].Coeffs, sk1.Coeffs) {
		t.Errorf("The first shared polynomial is not equal expected!")
	}

	if !reflect.DeepEqual(sk.Vec[1].Coeffs, sk2.Coeffs) {
		t.Errorf("The second shared polynomial is not equal expected!")
	}
}
