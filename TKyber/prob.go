package TKyber

import (
	"math"
	"math/rand"

	"ThresholdKyber.com/m/kyber"
)

func sampleRoundedGaussianPoly(q int, deg int, sigma_flood int) *kyber.Poly {
	coeffs_unrounded := make([]float64, deg+1)
	for i := 0; i < deg+1; i++ {
		coeffs_unrounded[i] = rand.NormFloat64() * float64(sigma_flood)
	}

	coeffs := make([]uint16, deg+1)
	for i := 0; i < deg+1; i++ {
		coeffs[i] = uint16(math.Round(coeffs_unrounded[i]))
	}

	out := new(kyber.Poly)
	copy(out.Coeffs[:], coeffs)

	return out
}
