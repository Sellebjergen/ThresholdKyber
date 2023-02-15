package owcpa_TKyber

import (
	"reflect"
	"testing"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
	"ThresholdKyber.com/m/util"
)

// ================= Share tests =================
func TestSharePolynomial(t *testing.T) {
	var toShare kyberk2so.Poly
	toShare[0] = 1
	toShare[1] = 2
	toShare[2] = 3
	toShare[3] = 4
	toShare[4] = 5

	shares := SharePolynomial(toShare, 10)

	var recombined kyberk2so.Poly
	for _, share := range shares {
		recombined = kyberk2so.PolyAdd(recombined, share)
	}
	recombined = kyberk2so.PolyReduce(recombined)

	if !reflect.DeepEqual(recombined, toShare) {
		t.Errorf("Recombined is not equal the original shared polynomial!")
	}
}

// ================= Share then Rec polynomial test =================
func TestShareThenRecPolynomial(t *testing.T) {
	lss := &LSSAdditive{}
	var toShare kyberk2so.Poly
	toShare[0] = 1
	toShare[1] = 2
	toShare[2] = 3
	toShare[3] = 4
	toShare[4] = 5

	polyShares := SharePolynomial(toShare, 10)

	// TODO: I had to transpose unfortunately
	shares := util.Transpose([][]kyberk2so.Poly{polyShares})

	recombined := lss.Rec(shares)

	if !reflect.DeepEqual(recombined, toShare) {
		t.Errorf("Recombined is not equal the original shared polynomial!")
	}
}

// ================= Testing that Share on polynomial ring ====
func TestShare(t *testing.T) {
	sk := kyberk2so.PolyvecNew(kyberk2so.ParamsK)
	lss := &LSSAdditive{}
	sk[0][0] = 42
	sk[0][1] = 73
	sk[1][0] = 0
	sk[1][1] = 27
	shares := lss.Share(sk, 3)

	// combine the first polynomial
	var sk1 kyberk2so.Poly
	sk1 = kyberk2so.PolyAdd(sk1, shares[0][0][0])
	sk1 = kyberk2so.PolyAdd(sk1, shares[1][0][0])
	sk1 = kyberk2so.PolyAdd(sk1, shares[2][0][0])

	// combine the second polynomial
	var sk2 kyberk2so.Poly
	sk2 = kyberk2so.PolyAdd(sk2, shares[0][0][1])
	sk2 = kyberk2so.PolyAdd(sk2, shares[1][0][1])
	sk2 = kyberk2so.PolyAdd(sk2, shares[2][0][1])

	kyberk2so.PolyReduce(sk1)
	kyberk2so.PolyReduce(sk2)

	if !reflect.DeepEqual(sk[0], sk1) {
		t.Errorf("The first shared polynomial is not equal expected!")
	}

	if !reflect.DeepEqual(sk[1], sk2) {
		t.Errorf("The second shared polynomial is not equal expected!")
	}
}
