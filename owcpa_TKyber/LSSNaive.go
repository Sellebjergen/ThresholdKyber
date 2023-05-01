package owcpa_TKyber

import (
	"fmt"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
	"ThresholdKyber.com/m/util"
)

type LSSNaive struct{}

func (s *LSSNaive) Share(sk kyberk2so.PolyVec, n int, t int) [][]kyberk2so.PolyVec {
	r := len(sk)

	combinations := util.MakeCombinations(n, t+1)

	fmt.Println(combinations)

	skShares := make([][][]kyberk2so.Poly, len(combinations))
	for i := 0; i < len(combinations); i++ {
		skShares[i] = make([][]kyberk2so.Poly, r)
		for j, e := range sk {
			skShares[i][j] = SharePolynomial(e, t+1)
		}
	}

	// the total amount of shares
	shares := make([][]kyberk2so.PolyVec, n)
	for i := 0; i < n; i++ {
		shares[i] = make([]kyberk2so.PolyVec, len(combinations))
		for j := 0; j < len(combinations); j++ {
			shares[i][j] = kyberk2so.PolyvecNew(r)
		}
	}

	for indx, comb := range combinations {
		for share_num, player := range comb {
			for poly_num := 0; poly_num < r; poly_num++ {
				copy(shares[player-1][indx][poly_num][:], skShares[indx][poly_num][share_num][:])
			}
		}
	}

	fmt.Println("SHARES")
	fmt.Println(shares)

	return shares
}

func (s *LSSNaive) Rec(d_is [][]kyberk2so.Poly, n int, t int) kyberk2so.Poly {
	var p kyberk2so.Poly

	combinations := util.MakeCombinations(n, t+1)

	fmt.Println("YAAAA")
	fmt.Println(d_is)

	for j := 0; j < len(combinations); j++ {
		comb := combinations[j]
		added := 0
		for i := 0; i < n; i++ {

			hasShare := util.Contains(comb, i+1)
			if hasShare { // TODO: Do we need  && (d_is[i][j] != kyberk2so.Poly{0}) ???
				added += 1
				p = kyberk2so.PolyAdd(p, d_is[i][j])
				if added == t+1 {
					return kyberk2so.PolyReduce(p)
				}
			}
		}
	}

	return kyberk2so.PolyReduce(p)
}
