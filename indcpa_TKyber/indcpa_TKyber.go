package indcpa_TKyber

import (
	"reflect"

	"ThresholdKyber.com/m/kyber"
	owcpa "ThresholdKyber.com/m/owcpa_TKyber"
)

type indcpaCiphertext struct {
	cF         *kyber.Poly
	encyptions [][]byte
	cG         []byte
}

func Setup(params *owcpa.OwcpaParams, n int, t int) (*kyber.IndcpaPublicKey, []kyber.PolyVec) {
	return owcpa.Setup(params, n, t)
}

func Enc(params *owcpa.OwcpaParams, msg []byte, pk *kyber.IndcpaPublicKey, delta int) *indcpaCiphertext {
	var mp kyber.Poly
	mp.ToMsg(msg)
	x := make([]*kyber.Poly, delta)
	for i := 0; i < delta; i++ {
		x[i] = owcpa.SampleUnifPolynomial(2)
	}
	c := new(indcpaCiphertext)
	c.encyptions = make([][]byte, delta)

	c.cF.Add(&mp, F(x))
	for i := 0; i < delta; i++ {
		xi_bytes := make([]byte, 32)
		x[i].ToMsg(xi_bytes)
		c.encyptions[i] = owcpa.Enc(params, xi_bytes, pk)
	}
	c.cG = G(x)

	return c
}

func PartDec(params *owcpa.OwcpaParams, sk_i kyber.PolyVec, ct indcpaCiphertext, party int, delta int) []*kyber.Poly {
	d_i := make([]*kyber.Poly, delta)
	for j := 0; j < delta; j++ {
		c_j := ct.encyptions[j]
		d_i[j] = owcpa.PartDec(params, sk_i, c_j, party)
	}
	return d_i
}

func Combine(params *owcpa.OwcpaParams, ct indcpaCiphertext, d_is [][]*kyber.Poly) *kyber.Poly {
	delta := len(ct.encyptions)
	var mp kyber.Poly
	x_prime := make([]*kyber.Poly, delta)
	for j := 0; j < delta; j++ {
		x_prime[j] = owcpa.Combine(params, ct.encyptions[j], d_is[j])
	}
	mp.Sub(ct.cF, F(x_prime))

	if !reflect.DeepEqual(ct.cG, G(x_prime)) {
		return nil
	}

	return &mp
}

func F(x []*kyber.Poly) *kyber.Poly {
	return nil
}

func G(x []*kyber.Poly) []byte {
	return nil
}
