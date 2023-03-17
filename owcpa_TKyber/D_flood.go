package owcpa_TKyber

import (
	"crypto/rand"
	"math"
	mrand "math/rand"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
	"ThresholdKyber.com/m/util"
)

type GaussianNoiseDist struct {
	SigmaFlood int
}

type BinomialNoiseDist struct {
	eta int
}

// TODO: UNSAFE, does not use crypto/rand
func (d *GaussianNoiseDist) SampleNoise(params *OwcpaParams, deg int) kyberk2so.Poly {
	coeffs_unrounded := make([]float64, deg+1)
	for i := 0; i < deg+1; i++ {
		coeffs_unrounded[i] = mrand.NormFloat64() * float64(d.SigmaFlood)
	}

	coeffs := make([]int16, deg+1)
	for i := 0; i < deg+1; i++ {
		coeffs[i] = int16(math.Round(coeffs_unrounded[i]))
	}

	var out kyberk2so.Poly
	copy(out[:], coeffs)

	return out
}

// Implemented as in Kyber specification 3.2
// Uses crypto/rand, so sampling is cryptographically secure
func (d *BinomialNoiseDist) SampleNoise(params *OwcpaParams, deg int) kyberk2so.Poly {
	b := make([]byte, 64*d.eta)
	rand.Read(b)

	var f kyberk2so.Poly
	beta := util.BytesToBits(b)
	for i := 0; i < 256; i++ {
		a := int16(0)
		b := int16(0)
		for j := 0; j < d.eta; j++ {
			a += int16(beta[2*i*d.eta+j])
			b += int16(beta[2*i*d.eta+d.eta+j])
		}
		f[i] = a - b
	}
	return f
}
