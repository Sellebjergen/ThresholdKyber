package owcpa_TKyber

import (
	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
	"reflect"
	"testing"
)

func TestShareGivesCorrectAmountOfShares(t *testing.T) {
	amountOfPlayers := 3
	MockKyber512Key := SampleUniformPolyVec(17, 2)

	shares := RepShare(MockKyber512Key, amountOfPlayers)

	if !reflect.DeepEqual(len(shares), amountOfPlayers) {
		t.Errorf("Something went wrong, combination of shares are not equal starting value!")
	}
}

func TestSharesCanBeReconstructed(t *testing.T) {
	amountOfPlayers := 3
	MockKyber512Key := SampleUniformPolyVec(17, 2)

	shares := RepShare(MockKyber512Key, amountOfPlayers)

	var p1 kyberk2so.Poly
	p1 = kyberk2so.PolyAdd(p1, shares[0][0][1]) // player 1 share of first poly, Share 2
	p1 = kyberk2so.PolyAdd(p1, shares[0][0][2]) // player 1 share of first poly, Share 3
	p1 = kyberk2so.PolyAdd(p1, shares[1][0][0]) // player 2 share of first poly, Share 1

	var p2 kyberk2so.Poly
	p2 = kyberk2so.PolyAdd(p2, shares[0][1][1]) // player 1 share of second poly. Share 2
	p2 = kyberk2so.PolyAdd(p2, shares[0][1][2]) // player 1 share of second poly. Share 3
	p2 = kyberk2so.PolyAdd(p2, shares[2][1][0]) // player 3 share of second poly. Share 1

	reconstructed := kyberk2so.PolyVec{p1, p2}

	if !reflect.DeepEqual(MockKyber512Key, reconstructed) {
		t.Errorf("Something went wrong assembling the mock key!")
	}
}

func TestSharesCanBeReconstructedUsingARealKyberKey(t *testing.T) {
	amountOfPlayers := 3
	sk, _, _ := kyberk2so.IndcpaKeypair(2)
	unpackedSk := kyberk2so.IndcpaUnpackPrivateKey(sk, 2)

	shares := RepShare(unpackedSk, amountOfPlayers)

	// todo: this seems like the nesting is to deep? shouldn't this only be a matrix?
	var p1 kyberk2so.Poly
	// take 2 shares from player 1 and the last share he does not have from player 2
	p1 = kyberk2so.PolyAdd(p1, shares[0][0][1]) // player 1 share of first poly, Share 2
	p1 = kyberk2so.PolyAdd(p1, shares[0][0][2]) // player 1 share of first poly, Share 3
	p1 = kyberk2so.PolyAdd(p1, shares[1][0][0]) // player 2 share of first poly, Share 1

	var p2 kyberk2so.Poly
	p2 = kyberk2so.PolyAdd(p2, shares[0][1][1]) // player 1 share of second poly. Share 2
	p2 = kyberk2so.PolyAdd(p2, shares[0][1][2]) // player 1 share of second poly. Share 3
	p2 = kyberk2so.PolyAdd(p2, shares[2][1][0]) // player 3 share of second poly. Share 1

	reconstructed := kyberk2so.PolyVec{p1, p2}

	if !reflect.DeepEqual(unpackedSk, reconstructed) {
		t.Errorf("Something went wrong assembling the mock key!")
	}
}
