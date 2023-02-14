package owcpa_TKyber

import (
	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
)

func RepShare(sk kyberk2so.PolyVec, n int) [][]kyberk2so.PolyVec {
	skShares := make([]kyberk2so.PolyVec, len(sk))

	for i, e := range sk {
		skShares[i] = SharePolynomial(e, n)
	}

	// the total amount of shares
	shares := make([][]kyberk2so.PolyVec, n)

	// player 1 should have
	shares10 := kyberk2so.PolyVec{
		kyberk2so.Poly{},
		skShares[0][1],
		skShares[0][2],
	}
	shares11 := kyberk2so.PolyVec{
		kyberk2so.Poly{},
		skShares[1][1],
		skShares[1][2],
	}
	shares1 := []kyberk2so.PolyVec{
		shares10,
		shares11,
	}
	shares[0] = shares1

	// player 2 should have
	shares20 := kyberk2so.PolyVec{
		skShares[0][0],
		kyberk2so.Poly{},
		skShares[0][2],
	}
	shares21 := kyberk2so.PolyVec{
		skShares[1][0],
		kyberk2so.Poly{},
		skShares[1][2],
	}
	shares2 := []kyberk2so.PolyVec{
		shares20,
		shares21,
	}
	shares[1] = shares2

	// player 3 should have
	shares30 := kyberk2so.PolyVec{
		skShares[0][0],
		skShares[0][1],
		kyberk2so.Poly{},
	}
	shares31 := kyberk2so.PolyVec{
		skShares[1][0],
		skShares[1][1],
		kyberk2so.Poly{},
	}
	shares3 := []kyberk2so.PolyVec{
		shares30,
		shares31,
	}
	shares[2] = shares3

	// combines the different shares
	combinedShares := make([][]kyberk2so.PolyVec, n)
	combinedShares[0] = shares1
	combinedShares[1] = shares2
	combinedShares[2] = shares3

	return combinedShares
}

func RepRec(d_is []kyberk2so.Poly) kyberk2so.Poly {
	return *new(kyberk2so.Poly)
}
