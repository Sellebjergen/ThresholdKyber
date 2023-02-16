package owcpa_TKyber

import (
	"fmt"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
	"ThresholdKyber.com/m/util"
)

type LSSReplicated struct{}

func (s *LSSReplicated) Share(sk kyberk2so.PolyVec, n int, t int) [][]kyberk2so.PolyVec {
	return ShareRepNaive(sk, n, t, false)
}

func (s *LSSReplicated) Rec(d_is [][]kyberk2so.Poly, n int, t int) kyberk2so.Poly {
	return RecRepNaive(d_is, n, t, false)
}

func ShareRepNaive(sk kyberk2so.PolyVec, n int, t int, isNaive bool) [][]kyberk2so.PolyVec {
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
			shouldGetShare := util.Contains(comb, i+1)
			if !isNaive {
				shouldGetShare = !shouldGetShare
			}
			if shouldGetShare {
				// Iterate over the r = k polynomials of the sk
				for poly_num := 0; poly_num < r; poly_num++ {
					copy(shares[i][j][poly_num][:], skShares[poly_num][j][:])
				}

			}

		}

	}

	//fmt.Println(shares)
	return shares
}

func RecRepNaive(d_is [][]kyberk2so.Poly, n int, t int, isNaive bool) kyberk2so.Poly {
	//fmt.Println(d_is)
	var p1 kyberk2so.Poly

	combinations := util.MakeCombinations(n, t)

	fmt.Println(combinations)

	for j := 0; j < len(combinations); j++ {
		comb := combinations[j]
		for i := 0; i < n; i++ {
			hasShare := util.Contains(comb, i+1)
			if !isNaive {
				hasShare = !hasShare
			}
			if hasShare {
				p1 = kyberk2so.PolyAdd(p1, d_is[i][j])
				break

			}
		}
	}

	return kyberk2so.PolyReduce(p1)
}
