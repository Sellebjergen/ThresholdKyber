package TKyber

type polyRing struct {
	q   int
	mod *Polynomial
}

type Polynomial struct {
	Coeffs []int32
}

func (*polyRing) init() *polyRing {
	rq := &polyRing{
		q:   3329,
		mod: getModulusPoly(),
	}

	return rq
}

func (r *polyRing) add(a, b *Polynomial) *Polynomial {
	out := make([]int32, max(len(a.Coeffs), len(b.Coeffs)))
	for i, coef := range a.Coeffs {
		out[i] += coef
	}
	for i, coef := range b.Coeffs {
		out[i] += coef
	}

	pre_reduce := Polynomial{Coeffs: out}
	return r.reduce(pre_reduce)
}

func (r *polyRing) sub(a, b *Polynomial) *Polynomial {
	out := make([]int32, max(len(a.Coeffs), len(b.Coeffs)))
	for i, coef := range a.Coeffs {
		out[i] += coef
	}
	for i, coef := range b.Coeffs {
		out[i] -= coef
	}

	pre_reduce := Polynomial{Coeffs: out}
	return r.reduce(pre_reduce)
}

func (r *polyRing) mult(a, b *Polynomial) *Polynomial {
	out := make([]int32, len(a.Coeffs)+len(b.Coeffs)-1)

	for i := 0; i < len(a.Coeffs); i++ {
		for j := 0; j < len(b.Coeffs); j++ {
			out[i+j] += a.Coeffs[i] * b.Coeffs[j]
		}
	}

	pre_reduce := Polynomial{Coeffs: out}

	return r.reduce(pre_reduce)
}

func (r *polyRing) mult_const(a *Polynomial, c int32) *Polynomial {
	out := make([]int32, len(a.Coeffs))
	for i := 0; i < len(a.Coeffs); i++ {
		out[i] = a.Coeffs[i] * c
	}
	pre_reduce := Polynomial{Coeffs: out}

	return r.reduce(pre_reduce)
}

func (r *polyRing) polynomialLongDivision(pol Polynomial) {
	// TODO
}

func (r *polyRing) rem_syntheticLongDivison(pol Polynomial) *Polynomial {
	if r.mod.getDeg() > pol.getDeg() {
		return &pol
	}

	out := Reverse(pol.Coeffs)
	divisor := Reverse(r.mod.Coeffs)
	normalizer := divisor[0]
	for i := 0; i < pol.getDeg()-r.mod.getDeg()+1; i++ {
		out[i] /= normalizer
		coef := out[i]
		if coef != 0 {
			for j := 1; j < len(divisor); j++ {
				out[i+j] += -divisor[j] * coef
			}
		}
	}
	final := Reverse(out)
	return &Polynomial{Coeffs: final[:len(divisor)-1]}
}

func (r *polyRing) reduce(pol Polynomial) *Polynomial {
	rem := r.rem_syntheticLongDivison(pol)
	out := rem
	// Compute mod q for each coeff
	for i := 0; i < len(out.Coeffs); i++ {
		out.Coeffs[i] = mod(out.Coeffs[i], int32(r.q)) // TODO: NOT EUCLIDEAN MODULO
	}
	return trimPoly(out)
}

func getModulusPoly() *Polynomial {
	res := make([]int32, 256)
	res[0] = 1
	res[255] = 1

	return &Polynomial{Coeffs: res}
}

/* func (p *Polynomial) toKyberPoly() *kyber.Poly {
	return &kyber.Poly{Coeffs: *(*[256]uint16)(p.Coeffs)}
} */

func (p *Polynomial) getDeg() int {
	return len(p.Coeffs)
}

func trimPoly(p *Polynomial) *Polynomial {
	coeffs := p.Coeffs
	for coeffs[len(coeffs)-1] == 0 {
		coeffs = coeffs[:len(coeffs)-1]
	}
	return &Polynomial{Coeffs: coeffs}
}

func (p *Polynomial) Copy() *Polynomial {
	out_coef := make([]int32, len(p.Coeffs))
	copy(out_coef, p.Coeffs)
	return &Polynomial{Coeffs: out_coef}
}

func (p *Polynomial) lead() int32 {
	return p.Coeffs[len(p.Coeffs)-1]
}
