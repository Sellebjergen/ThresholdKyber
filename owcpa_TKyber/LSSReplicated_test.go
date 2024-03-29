package owcpa_TKyber

import (
	"reflect"
	"testing"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
)

func TestShareGivesCorrectAmountOfShares(t *testing.T) {
	amountOfPlayers := 3
	MockKyber512Key := SampleUniformPolyVec(17, 2)

	shares := new(LSSReplicated).Share(MockKyber512Key, amountOfPlayers, 1)

	if !reflect.DeepEqual(len(shares), amountOfPlayers) {
		t.Errorf("Something went wrong, combination of shares are not equal starting value!")
	}
}

func TestSharesCanBeReconstructed(t *testing.T) {
	amountOfPlayers := 3
	MockKyber512Key := SampleUniformPolyVec(17, 2)

	shares := new(LSSReplicated).Share(MockKyber512Key, amountOfPlayers, 1)

	var p1 kyberk2so.Poly
	p1 = kyberk2so.PolyAdd(p1, shares[0][1][0]) // player 1 share of first poly, Share 2
	p1 = kyberk2so.PolyAdd(p1, shares[0][2][0]) // player 1 share of first poly, Share 3
	p1 = kyberk2so.PolyAdd(p1, shares[1][0][0]) // player 2 share of first poly, Share 1

	var p2 kyberk2so.Poly
	p2 = kyberk2so.PolyAdd(p2, shares[0][1][1]) // player 1 share of second poly. Share 2
	p2 = kyberk2so.PolyAdd(p2, shares[0][2][1]) // player 1 share of second poly. Share 3
	p2 = kyberk2so.PolyAdd(p2, shares[2][0][1]) // player 3 share of second poly. Share 1

	reconstructed := kyberk2so.PolyVec{p1, p2}

	if !reflect.DeepEqual(MockKyber512Key, reconstructed) {
		t.Errorf("Something went wrong assembling the mock key!")
	}
}

func TestSharesCanBeReconstructedUsingARealKyberKey(t *testing.T) {
	amountOfPlayers := 3
	sk, _, _ := kyberk2so.IndcpaKeypair(2)
	unpackedSk := kyberk2so.IndcpaUnpackPrivateKey(sk, 2)

	shares := new(LSSReplicated).Share(unpackedSk, amountOfPlayers, 1)

	// todo: this seems like the nesting is to deep? shouldn't this only be a matrix?
	var p1 kyberk2so.Poly
	// take 2 shares from player 1 and the last share he does not have from player 2
	p1 = kyberk2so.PolyAdd(p1, shares[0][1][0]) // player 1 share of first poly, Share 2
	p1 = kyberk2so.PolyAdd(p1, shares[0][2][0]) // player 1 share of first poly, Share 3
	p1 = kyberk2so.PolyAdd(p1, shares[1][0][0]) // player 2 share of first poly, Share 1

	var p2 kyberk2so.Poly
	p2 = kyberk2so.PolyAdd(p2, shares[0][1][1]) // player 1 share of second poly. Share 2
	p2 = kyberk2so.PolyAdd(p2, shares[0][2][1]) // player 1 share of second poly. Share 3
	p2 = kyberk2so.PolyAdd(p2, shares[2][0][1]) // player 3 share of second poly. Share 1

	reconstructed := kyberk2so.PolyVec{p1, p2}

	if !reflect.DeepEqual(unpackedSk, reconstructed) {
		t.Errorf("Something went wrong assembling the mock key!")
	}
}

func TestSinglePolyCanBeReconstructed(t *testing.T) {
	toShare := []kyberk2so.Poly{{1, 2, 3, 4, 5, 6}}
	shared := new(LSSReplicated).Share(toShare, 3, 1)

	var p1 kyberk2so.Poly
	p1 = kyberk2so.PolyAdd(p1, shared[0][1][0])
	p1 = kyberk2so.PolyAdd(p1, shared[0][2][0])
	p1 = kyberk2so.PolyAdd(p1, shared[1][0][0])

	if !reflect.DeepEqual(toShare[0], p1) {
		t.Errorf("WEE WOO WEE WOO")
	}

	var zero kyberk2so.Poly

	if !reflect.DeepEqual(shared[0][0][0], zero) {
		t.Errorf("WEE WOO WEE WOO")
	}

	if !reflect.DeepEqual(shared[1][1][0], zero) {
		t.Errorf("WEE WOO WEE WOO")
	}

	if !reflect.DeepEqual(shared[2][2][0], zero) {
		t.Errorf("WEE WOO WEE WOO")
	}
}

func TestSinglePolyCanBeReconstructedN4T2(t *testing.T) {
	toShare := []kyberk2so.Poly{{1, 2, 3, 4, 5, 6}}
	shared := new(LSSReplicated).Share(toShare, 4, 2)

	var p1 kyberk2so.Poly
	p1 = kyberk2so.PolyAdd(p1, shared[0][3][0])
	p1 = kyberk2so.PolyAdd(p1, shared[0][4][0])
	p1 = kyberk2so.PolyAdd(p1, shared[0][5][0])
	p1 = kyberk2so.PolyAdd(p1, shared[2][0][0])
	p1 = kyberk2so.PolyAdd(p1, shared[1][1][0])
	p1 = kyberk2so.PolyAdd(p1, shared[1][2][0])

	if !reflect.DeepEqual(toShare[0], p1) {
		t.Errorf("WEE WOO WEE WOO")
	}
}

func TestReplicatedRecSimple(t *testing.T) {
	one := kyberk2so.Poly{1}
	zero := kyberk2so.Poly{0}
	expected := kyberk2so.Poly{3}
	d_is := [][]kyberk2so.Poly{}
	d_is = append(d_is, []kyberk2so.Poly{zero, one, one})
	d_is = append(d_is, []kyberk2so.Poly{one, zero, one})
	d_is = append(d_is, []kyberk2so.Poly{one, one, zero})

	lss := &LSSReplicated{}
	res := lss.Rec(d_is, 3, 1)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("WEE WOO WEE WOO")
	}
}

func TestReplicatedRecAdvanced(t *testing.T) {
	one := kyberk2so.Poly{1}
	two := kyberk2so.Poly{2}
	five := kyberk2so.Poly{5}
	zero := kyberk2so.Poly{0}
	expected := kyberk2so.Poly{8}
	d_is := [][]kyberk2so.Poly{}
	d_is = append(d_is, []kyberk2so.Poly{zero, two, five})
	d_is = append(d_is, []kyberk2so.Poly{one, zero, five})
	d_is = append(d_is, []kyberk2so.Poly{one, two, zero})

	lss := &LSSReplicated{}
	res := lss.Rec(d_is, 3, 1)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("WEE WOO WEE WOO")
	}
}

// Share test
func TestReplicatedSharesAreEqual(t *testing.T) {
	toShare := []kyberk2so.Poly{{1, 2, 3, 4, 5, 6}}
	shared := new(LSSReplicated).Share(toShare, 3, 1)

	if !reflect.DeepEqual(shared[0][1], shared[2][1]) {
		t.Errorf("WEE WOO WEE WOO")
	}

	if !reflect.DeepEqual(shared[2][0], shared[1][0]) {
		t.Errorf("WEE WOO WEE WOO")
	}

	if !reflect.DeepEqual(shared[0][2], shared[1][2]) {
		t.Errorf("WEE WOO WEE WOO")
	}
}
