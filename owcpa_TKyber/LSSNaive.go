package owcpa_TKyber

import (
	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
)

type LSSNaive struct{}

func (s *LSSNaive) Share(sk kyberk2so.PolyVec, n int, t int) [][]kyberk2so.PolyVec {
	return ShareRepNaive(sk, n, t+1, true)
}

func (s *LSSNaive) Rec(d_is [][]kyberk2so.Poly, n int, t int) kyberk2so.Poly {
	return RecRepNaive(d_is, n, t+1, true)
}
