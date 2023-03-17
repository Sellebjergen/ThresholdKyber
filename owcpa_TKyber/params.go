package owcpa_TKyber

import (
	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
)

type OwcpaParams struct {
	Ell          int
	Q            int
	D_flood_dist NoiseDistribution
	LSS_scheme   LSSScheme
}

type NoiseDistribution interface {
	SampleNoise(params *OwcpaParams, deg int) kyberk2so.Poly
}

type LSSScheme interface {
	Share(sk kyberk2so.PolyVec, n int, t int) [][]kyberk2so.PolyVec
	Rec(d_is [][]kyberk2so.Poly, n int, t int) kyberk2so.Poly
}

func NewParameterSet(name string) *OwcpaParams {
	var p OwcpaParams
	switch name {
	case "TKyber1024-Q16645":
		p.Ell = 1
		p.D_flood_dist = &GaussianNoiseDist{SigmaFlood: 947}
		p.LSS_scheme = &LSSAdditive{}
	case "TKyber1024-Q33290":
		p.Ell = 2
		p.Q = 3329 * 10
		p.D_flood_dist = &GaussianNoiseDist{SigmaFlood: 1994}
		p.LSS_scheme = &LSSAdditive{}
	case "TKyber1024-Q29961":
		p.Ell = 1
		p.Q = 3329 * 9
		p.D_flood_dist = &GaussianNoiseDist{SigmaFlood: 1197}
		p.LSS_scheme = &LSSAdditive{}
	case "TKyber-Test":
		p.Ell = 1
		p.Q = 3329
		p.D_flood_dist = &GaussianNoiseDist{SigmaFlood: 75}
		p.LSS_scheme = &LSSAdditive{}
	case "TKyber-Test-Replicated":
		p.Ell = 1
		p.Q = 3329
		p.D_flood_dist = &GaussianNoiseDist{SigmaFlood: 75}
		p.LSS_scheme = &LSSReplicated{}
	case "TKyber-Test-Naive":
		p.Ell = 1
		p.Q = 3329
		p.D_flood_dist = &GaussianNoiseDist{SigmaFlood: 75}
		p.LSS_scheme = &LSSNaive{}
	default:
		panic("Error: Name did not match existing parameter set")
	}

	return &p
}
