package owcpa_TKyber

import (
	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
	"ThresholdKyber.com/m/util"
)

type LSSReplicated struct{}

func (s *LSSReplicated) Share(sk kyberk2so.PolyVec, n int, t int) [][]kyberk2so.PolyVec {
	r := len(sk)
	skShares := make([][]kyberk2so.Poly, r)

	combinations := util.MakeCombinations(n, t)

	// the total amount of shares
	shares := make([][]kyberk2so.PolyVec, n)
	for i := 0; i < n; i++ {
		shares[i] = make([]kyberk2so.PolyVec, len(combinations))
		for j := 0; j < len(combinations); j++ {
			shares[i][j] = kyberk2so.PolyvecNew(r)
		}
	}

	for i, e := range sk {
		skShares[i] = SharePolynomial(e, len(combinations))
	}

	// Players
	for i := 0; i < n; i++ {
		// Combinations
		for j := 0; j < len(combinations); j++ {
			comb := combinations[j]
			shouldGetShare := !util.Contains(comb, i+1)
			if shouldGetShare {
				// Iterate over the r = k polynomials of the sk
				for poly_num := 0; poly_num < r; poly_num++ {
					copy(shares[i][j][poly_num][:], skShares[poly_num][j][:])
				}

			}

		}

	}

	return shares
}

func (s *LSSReplicated) Rec(d_is [][]kyberk2so.Poly, n int, t int) kyberk2so.Poly {
	var p kyberk2so.Poly

	combinations := util.MakeCombinations(n, t)

	for j := 0; j < len(combinations); j++ {
		comb := combinations[j]
		for i := 0; i < n; i++ {
			hasShare := !util.Contains(comb, i+1)
			if hasShare { // TODO: Do we need  && (d_is[i][j] != kyberk2so.Poly{0}) ???
				p = kyberk2so.PolyAdd(p, d_is[i][j])
				break

			}
		}
	}

	return kyberk2so.PolyReduce(p)
}
