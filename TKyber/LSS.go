package TKyber

import (
	"ThresholdKyber.com/m/kyber"
)

type share struct {
	poly *Polynomial
}

// Represents additively secret sharing
func Share(params kyber.ParameterSet, sk *kyber.IndcpaSecretKey) []*share {
	polyVec_sk := params.AllocPolyVec()
	kyber.UnpackSecretKey(&polyVec_sk, sk.Packed)

	return nil
}

func (r *polyRing) Rec(d_is []*share) *Polynomial {
	out := d_is[0].poly

	for i := 1; i < len(d_is); i++ {
		out = r.add(out, d_is[i].poly)
	}

	return out
}
