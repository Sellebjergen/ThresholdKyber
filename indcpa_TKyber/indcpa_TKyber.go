package indcpa_TKyber

import (
	"reflect"

	"golang.org/x/crypto/sha3"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
	owcpa "ThresholdKyber.com/m/owcpa_TKyber"
	"ThresholdKyber.com/m/util"
)

type indcpaCiphertext struct {
	cF         kyberk2so.Poly
	encyptions [][]byte
	cG         []byte
}

func Setup(params *owcpa.OwcpaParams, n int, t int) ([]byte, [][]kyberk2so.PolyVec) {
	return owcpa.Setup(params, n, t)
}

func Enc(params *owcpa.OwcpaParams, msg []byte, pk []byte, delta int) *indcpaCiphertext {
	mp := kyberk2so.PolyFromMsg(msg)
	x := make([]kyberk2so.Poly, delta)
	for i := 0; i < delta; i++ {
		x[i] = owcpa.SampleUnifPolynomial(2)
	}

	c := new(indcpaCiphertext)
	c.encyptions = make([][]byte, delta)

	c0 := kyberk2so.PolyAdd(mp, F(x))
	c.cF = c0
	for i := 0; i < delta; i++ {
		upscaled := owcpa.Upscale(x[i], 2, params.Q)
		xi_bytes := kyberk2so.PolyToMsg(upscaled)
		c.encyptions[i] = owcpa.Enc(params, xi_bytes, pk)
	}
	c.cG = G(x)

	return c
}

func PartDec(params *owcpa.OwcpaParams, sk_i []kyberk2so.PolyVec, ct *indcpaCiphertext, party, delta int, n int, t int) [][]kyberk2so.Poly {
	d_i := make([][]kyberk2so.Poly, delta)
	for j := 0; j < delta; j++ {
		d_i[j] = owcpa.PartDec(params, sk_i, ct.encyptions[j], party, n, t)
	}
	return d_i
}

func Combine(params *owcpa.OwcpaParams, ct *indcpaCiphertext, d_is [][][]kyberk2so.Poly, n int, t int) kyberk2so.Poly {
	delta := len(ct.encyptions)

	x_prime := make([]kyberk2so.Poly, delta)

	// This is needed since originally 1 dim is player, second is which of the delta encryptions, third is L
	// We need to fetch all players shares of the i'th out of the delta encryptions
	// Therefore we transpose the first and second dimensions
	// To get new 3d matrix where 1 dim is delta encryptions, second is player, and third is L
	// TODO: This can probably be done cleaner
	d_is_transp := util.SwapFirstAndSecondDim(d_is)

	for j := 0; j < delta; j++ {
		combined := owcpa.Combine(params, ct.encyptions[j], d_is_transp[j], n, t)
		x_prime[j] = owcpa.Downscale(combined, 2, params.Q)
	}

	mp := kyberk2so.PolySub(ct.cF, F(x_prime))

	if !reflect.DeepEqual(ct.cG, G(x_prime)) {
		panic("Error: c_(delta + 1) != G(x')")
	}

	return mp
}

func F(x []kyberk2so.Poly) kyberk2so.Poly {
	hash := sha3.NewShake256()
	output := make([]byte, 384)
	for i := 0; i < len(x); i++ {
		poly_bytes := kyberk2so.PolyToBytes(x[i])
		hash.Write(poly_bytes)
	}
	hash.Read(output)
	return kyberk2so.PolyFromBytes(output)
}

func G(x []kyberk2so.Poly) []byte {
	hash := sha3.NewShake256()
	output := make([]byte, 2*100) // TODO: 100 er midlertidig, gider ikke på parameter safari lige nu
	for i := 0; i < len(x); i++ {
		poly_bytes := kyberk2so.PolyToBytes(x[i])
		hash.Write(poly_bytes)
	}
	hash.Read(output)
	return output
}
