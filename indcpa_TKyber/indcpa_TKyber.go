package indcpa_TKyber

import (
	"fmt"
	"reflect"

	"golang.org/x/crypto/sha3"

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
	mp.FromMsg(msg)
	x := make([]*kyber.Poly, delta)
	for i := 0; i < delta; i++ {
		x[i] = owcpa.SampleUnifPolynomial(2)
	}
	fmt.Println("x_1")
	fmt.Println(x[0])
	c := new(indcpaCiphertext)
	c.encyptions = make([][]byte, delta)

	mp.Add(&mp, F(x))
	c.cF = &mp
	for i := 0; i < delta; i++ {
		xi_bytes := make([]byte, 32)
		x[i].ToMsg(xi_bytes)
		c.encyptions[i] = owcpa.Enc(params, xi_bytes, pk)
	}
	c.cG = G(x)

	return c
}

func PartDec(params *owcpa.OwcpaParams, sk_i kyber.PolyVec, ct *indcpaCiphertext, party int, delta int) []*kyber.Poly {
	d_i := make([]*kyber.Poly, delta)
	for j := 0; j < delta; j++ {
		d_i[j] = owcpa.PartDec(params, sk_i, ct.encyptions[j], party)
	}
	return d_i
}

func Combine(params *owcpa.OwcpaParams, ct *indcpaCiphertext, d_is [][]*kyber.Poly) *kyber.Poly {
	delta := len(ct.encyptions)
	var mp kyber.Poly
	x_prime := make([]*kyber.Poly, delta)
	for j := 0; j < delta; j++ {
		x_prime[j] = owcpa.Combine(params, ct.encyptions[j], d_is[j])
	}
	fmt.Println("x_1'")
	fmt.Println(x_prime[0])
	mp.Sub(ct.cF, F(x_prime))

	if !reflect.DeepEqual(ct.cG, G(x_prime)) {
		panic("Error: c_(delta + 1) != G(x')")
	}

	return &mp
}

func F(x []*kyber.Poly) *kyber.Poly {
	hash := sha3.NewShake256()
	output := make([]byte, 13*(7681/8)+12) // TODO: Ved ikke om længden er korrekt
	for i := 0; i < len(x); i++ {
		poly_bytes := make([]byte, 13*(7681/8)+12)
		x[i].ToBytes(poly_bytes)
		hash.Write(poly_bytes)
	}
	hash.Read(output)

	p := new(kyber.Poly)
	p.FromBytes(output)
	return p
}

func G(x []*kyber.Poly) []byte {
	hash := sha3.NewShake256()
	output := make([]byte, 2*100) // TODO: 100 er midlertidig, gider ikke på parameter safari lige nu
	for i := 0; i < len(x); i++ {
		poly_bytes := make([]byte, 13*(7681/8)+12)
		x[i].ToBytes(poly_bytes)
		hash.Write(poly_bytes)
	}
	hash.Read(output)
	return output
}
