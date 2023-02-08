package owcpa_TKyber

import (
	"math"
	"math/rand"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
)

type GaussianNoiseDist struct{}

type BinomialNoiseDist struct{}

func (d *GaussianNoiseDist) SampleNoise(q int, deg int, sigma_flood int) kyberk2so.Poly {
	coeffs_unrounded := make([]float64, deg+1)
	for i := 0; i < deg+1; i++ {
		coeffs_unrounded[i] = rand.NormFloat64() * float64(sigma_flood)
	}

	coeffs := make([]int16, deg+1)
	for i := 0; i < deg+1; i++ {
		coeffs[i] = int16(math.Round(coeffs_unrounded[i]))
	}

	var out kyberk2so.Poly
	copy(out[:], coeffs)

	return out
}
