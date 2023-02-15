package owcpa_TKyber

import (
	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
)

type OwcpaParams struct {
	Ell          int
	Sigma        int
	Q            int
	D_flood_dist NoiseDistribution
	LSS_scheme   LSSScheme
}

type NoiseDistribution interface {
	SampleNoise(q int, deg int, sigma_flood int) kyberk2so.Poly
}

type LSSScheme interface {
	Share(sk kyberk2so.PolyVec, n int) [][]kyberk2so.PolyVec
	Rec(d_is [][]kyberk2so.Poly) kyberk2so.Poly
}

func NewParameterSet(name string) *OwcpaParams {
	var p OwcpaParams
	switch name {
	case "TKyber1024-Q16645":
		p.Ell = 1
		p.Sigma = 947 // This correct?
		p.Q = 3329 * 5
		p.D_flood_dist = &GaussianNoiseDist{}
		p.LSS_scheme = &LSSAdditive{}
	case "TKyber1024-Q33290":
		p.Ell = 2
		p.Sigma = 1994
		p.Q = 3329 * 10
		p.D_flood_dist = &GaussianNoiseDist{}
		p.LSS_scheme = &LSSAdditive{}
	case "TKyber1024-Q29961":
		p.Ell = 1
		p.Sigma = 1197
		p.Q = 3329 * 9
		p.D_flood_dist = &GaussianNoiseDist{}
		p.LSS_scheme = &LSSAdditive{}
	case "TKyber-Test":
		p.Ell = 1
		p.Sigma = 100
		p.Q = 3329
		p.D_flood_dist = &GaussianNoiseDist{}
		p.LSS_scheme = &LSSAdditive{}
	default:
		panic("Error: Name did not match existing parameter set")
	}

	return &p
}
