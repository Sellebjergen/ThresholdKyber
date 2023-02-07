package owcpa_TKyber

import "ThresholdKyber.com/m/kyber"

type OwcpaParams struct {
	KyberParams  *kyber.ParameterSet
	Ell          int
	Sigma        int
	Q            int
	D_flood_dist NoiseDistribution
}

type NoiseDistribution interface {
	SampleNoise(q int, deg int, sigma_flood int) *kyber.Poly
}

func newParameterSet(name string) *OwcpaParams {
	var p OwcpaParams
	switch name {
	case "TKyber1024-Q16645":
		p.KyberParams = kyber.Kyber1024
		p.Ell = 1
		p.Sigma = 947 // This correct?
		p.D_flood_dist = &GaussianNoiseDist{}
	case "TKyber1024-Q33290":
		p.KyberParams = kyber.Kyber1024
		p.Ell = 2
		p.Sigma = 1994
		p.D_flood_dist = &GaussianNoiseDist{}
	case "TKyber1024-Q29961":
		p.KyberParams = kyber.Kyber1024
		p.Ell = 1
		p.Sigma = 1197
		p.D_flood_dist = &GaussianNoiseDist{}
	case "TKyber-Test":
		p.KyberParams = kyber.Kyber512
		p.Ell = 1
		p.Sigma = 100
		p.D_flood_dist = &GaussianNoiseDist{}
	default:
		panic("Error: Name did not match existing parameter set")
	}

	return &p
}
