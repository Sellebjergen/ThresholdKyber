package TKyber

import (
	"math"
	"math/rand"

	"ThresholdKyber.com/m/kyber"
)

type Polynomial struct {
	Coeffs []int
}

func add(a, b *Polynomial) *Polynomial {
	out := make([]int, max(len(a.Coeffs), len(b.Coeffs)))
	for i, coef := range a.Coeffs {
		out[i] += coef
	}
	for i, coef := range b.Coeffs {
		out[i] += coef
	}

	return trimPoly(&Polynomial{Coeffs: out})
}

func sub(a, b *Polynomial) *Polynomial {
	out := make([]int, max(len(a.Coeffs), len(b.Coeffs)))
	for i, coef := range a.Coeffs {
		out[i] += coef
	}
	for i, coef := range b.Coeffs {
		out[i] -= coef
	}

	return trimPoly(&Polynomial{Coeffs: out})
}

func mult(a, b *Polynomial) *Polynomial {
	out := make([]int, len(a.Coeffs)+len(b.Coeffs)-1)

	for i := 0; i < len(a.Coeffs); i++ {
		for j := 0; j < len(b.Coeffs); j++ {
			out[i+j] += a.Coeffs[i] * b.Coeffs[j]
		}
	}

	return trimPoly(&Polynomial{Coeffs: out})
}

func mult_const(a *Polynomial, c int) *Polynomial {
	out := make([]int, len(a.Coeffs))
	for i := 0; i < len(a.Coeffs); i++ {
		out[i] = a.Coeffs[i] * c
	}
	return trimPoly(&Polynomial{Coeffs: out})
}

func neg(a *Polynomial) *Polynomial {
	return sub(&Polynomial{Coeffs: []int{0}}, a)
}

func (p *Polynomial) getDeg() int {
	return len(p.Coeffs)
}

func trimPoly(p *Polynomial) *Polynomial {
	coeffs := p.Coeffs
	for len(coeffs) > 0 && coeffs[len(coeffs)-1] == 0 {
		coeffs = coeffs[:len(coeffs)-1]
	}
	if len(coeffs) == 0 {
		return &Polynomial{Coeffs: []int{0}}
	}
	return &Polynomial{Coeffs: coeffs}
}

func (p *Polynomial) Copy() *Polynomial {
	out_coef := make([]int, len(p.Coeffs))
	copy(out_coef, p.Coeffs)
	return &Polynomial{Coeffs: out_coef}
}

func (p *Polynomial) toKyberPoly() *kyber.Poly {
	uint16_coeff := make([]uint16, len(p.Coeffs))
	for i, coef := range p.Coeffs {
		uint16_coeff[i] = uint16(coef)
	}

	var out_coeff [256]uint16

	copy(out_coeff[:], uint16_coeff)
	return &kyber.Poly{Coeffs: out_coeff}
}

func fromKyberPoly(p *kyber.Poly) *Polynomial {
	non_fixed_arr_coeff := p.Coeffs[:]
	new_coeff := make([]int, len(p.Coeffs))
	for i, coef := range non_fixed_arr_coeff {
		new_coeff[i] = int(coef)
	}
	return trimPoly(&Polynomial{Coeffs: new_coeff})
}

func samplePolyGaussian(q int, deg int, sigma_flood int) *Polynomial {
	coeffs_unrounded := make([]float64, deg+1)
	for i := 0; i < deg+1; i++ {
		coeffs_unrounded[i] = rand.NormFloat64() * float64(sigma_flood)
	}

	coeffs := make([]int, deg+1)
	for i := 0; i < deg+1; i++ {
		coeffs[i] = int(math.Round(coeffs_unrounded[i]))
	}

	return &Polynomial{Coeffs: coeffs}
}
