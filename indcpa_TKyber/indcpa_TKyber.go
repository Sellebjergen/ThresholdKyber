package indcpa_TKyber

import (
	"fmt"
	"reflect"

	"golang.org/x/crypto/sha3"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
	owcpa "ThresholdKyber.com/m/owcpa_TKyber"
)

type indcpaCiphertext struct {
	cF         kyberk2so.Poly
	encyptions [][]byte
	cG         []byte
}

func Setup(params *owcpa.OwcpaParams, n int, t int) ([]byte, []kyberk2so.PolyVec) {
	return owcpa.Setup(params, n, t)
}

func Enc(params *owcpa.OwcpaParams, msg []byte, pk []byte, delta int) *indcpaCiphertext {
	mp := kyberk2so.PolyFromMsg(msg)
	x := make([]kyberk2so.Poly, delta)
	for i := 0; i < delta; i++ {
		x[i] = owcpa.SampleUnifPolynomial(2)
	}
	fmt.Println("x_1")
	fmt.Println(x[0])
	c := new(indcpaCiphertext)
	c.encyptions = make([][]byte, delta)

	mp = kyberk2so.PolyAdd(mp, F(x))
	c.cF = mp
	for i := 0; i < delta; i++ {
		xi_bytes := kyberk2so.PolyToMsg(x[i])
		c.encyptions[i] = owcpa.Enc(params, xi_bytes, pk)
	}
	c.cG = G(x)

	return c
}

func PartDec(params *owcpa.OwcpaParams, sk_i kyberk2so.PolyVec, ct *indcpaCiphertext, party int, delta int) []kyberk2so.Poly {
	d_i := make([]kyberk2so.Poly, delta)
	for j := 0; j < delta; j++ {
		d_i[j] = owcpa.PartDec(params, sk_i, ct.encyptions[j], party)
	}
	return d_i
}

func Combine(params *owcpa.OwcpaParams, ct *indcpaCiphertext, d_is [][]kyberk2so.Poly) kyberk2so.Poly {
	delta := len(ct.encyptions)

	x_prime := make([]kyberk2so.Poly, delta)
	for j := 0; j < delta; j++ {
		x_prime[j] = owcpa.Combine(params, ct.encyptions[j], d_is[j])
	}
	fmt.Println("x_1'")
	fmt.Println(x_prime[0])
	mp := kyberk2so.PolySub(ct.cF, F(x_prime))

	if !reflect.DeepEqual(ct.cG, G(x_prime)) {
		panic("Error: c_(delta + 1) != G(x')")
	}

	return mp
}

func F(x []kyberk2so.Poly) kyberk2so.Poly {
	hash := sha3.NewShake256()
	output := make([]byte, 13*(7681/8)+12)
	for i := 0; i < len(x); i++ {
		poly_bytes := kyberk2so.PolyToBytes(x[i])
		hash.Write(poly_bytes)
	}
	hash.Read(output)
	return kyberk2so.PolyFromBytes(output) // TODO: Er message space R_2 for IND-CPA fint?
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
