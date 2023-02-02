package TKyber

type Polynomial struct {
	Coeffs []int32
}

func add(a, b *Polynomial) *Polynomial {
	out := make([]int32, max(len(a.Coeffs), len(b.Coeffs)))
	for i, coef := range a.Coeffs {
		out[i] += coef
	}
	for i, coef := range b.Coeffs {
		out[i] += coef
	}

	return trimPoly(&Polynomial{Coeffs: out})
}

func sub(a, b *Polynomial) *Polynomial {
	out := make([]int32, max(len(a.Coeffs), len(b.Coeffs)))
	for i, coef := range a.Coeffs {
		out[i] += coef
	}
	for i, coef := range b.Coeffs {
		out[i] -= coef
	}

	return trimPoly(&Polynomial{Coeffs: out})
}

func mult(a, b *Polynomial) *Polynomial {
	out := make([]int32, len(a.Coeffs)+len(b.Coeffs)-1)

	for i := 0; i < len(a.Coeffs); i++ {
		for j := 0; j < len(b.Coeffs); j++ {
			out[i+j] += a.Coeffs[i] * b.Coeffs[j]
		}
	}

	return trimPoly(&Polynomial{Coeffs: out})
}

func mult_const(a *Polynomial, c int32) *Polynomial {
	out := make([]int32, len(a.Coeffs))
	for i := 0; i < len(a.Coeffs); i++ {
		out[i] = a.Coeffs[i] * c
	}
	return trimPoly(&Polynomial{Coeffs: out})
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
		return &Polynomial{Coeffs: []int32{0}}
	}
	return &Polynomial{Coeffs: coeffs}
}

/* func (p *Polynomial) toKyberPoly() *kyber.Poly {
	return &kyber.Poly{Coeffs: *(*[256]uint16)(p.Coeffs)}
} */
