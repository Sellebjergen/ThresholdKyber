package TKyber

type quotRing struct {
	q   int
	mod *Polynomial
}

func (rq *quotRing) initKyberRing() *quotRing {
	rq.q = 3329
	rq.mod = getModulusPoly()

	return rq
}

func (r *quotRing) add(a, b *Polynomial) *Polynomial {
	pre_reduce := add(a, b)
	return r.reduce(pre_reduce)
}

func (r *quotRing) sub(a, b *Polynomial) *Polynomial {
	pre_reduce := sub(a, b)
	return r.reduce(pre_reduce)
}

func (r *quotRing) mult(a, b *Polynomial) *Polynomial {
	pre_reduce := mult(a, b)
	return r.reduce(pre_reduce)
}

func (r *quotRing) mult_const(a *Polynomial, c int) *Polynomial {
	pre_reduce := mult_const(a, c)
	return r.reduce(pre_reduce)
}

func (r *quotRing) neg(a *Polynomial) *Polynomial {
	pre_reduce := neg(a)
	return r.reduce(pre_reduce)
}

func (r *quotRing) polynomialLongDivision(pol Polynomial) {
	// TODO
}

func (r *quotRing) syntheticLongDivison(pol Polynomial) (*Polynomial, *Polynomial) {
	if r.mod.getDeg() > pol.getDeg() {
		return &Polynomial{Coeffs: []int{0}}, &pol
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
	return &Polynomial{Coeffs: final[len(divisor)-1:]}, &Polynomial{Coeffs: final[:len(divisor)-1]}
}

func (r *quotRing) reduce(pol *Polynomial) *Polynomial {
	_, rem := r.syntheticLongDivison(*pol)
	out := rem

	// Compute mod q for each coeff
	for i := 0; i < len(out.Coeffs); i++ {
		out.Coeffs[i] = euc_mod(out.Coeffs[i], int(r.q))
	}

	return trimPoly(out)
}

func getModulusPoly() *Polynomial {
	res := make([]int, 256)
	res[0] = 1
	res[255] = 1

	return &Polynomial{Coeffs: res}
}
