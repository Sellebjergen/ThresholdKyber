package gladius

import "fmt"

type PK struct {
	A1 [][]int64
	A2 [][]int64
}

type SK struct {
	R1 [][]int64
	Pk PK
}

type Ciphertext struct {
	C1 []int64
	C2 []int64
}

func keygen(n int, q int, ell int64) (PK, SK) {
	R1 := sampleRMatrix(n)
	R2 := sampleRMatrix(n)
	A1 := sampleUniform(n, q)
	G := constructGadgetMat(ell, n)

	A2 := matAdd(matAdd(matProd(A1, R1, n), R2), G)

	pk := PK{A1, A2}
	sk := SK{R1, pk}

	return pk, sk

}

func encrypt(n int, q int, ell int64, t int, p int64, msg []int64, pk PK) Ciphertext {
	c1 := vec_round_mod(vecMatProd(msg, pk.A1), p, q)
	c2 := vec_round_mod(vecMatProd(msg, pk.A2), p, q)
	return Ciphertext{c1, c2}
}

func decrypt(ct Ciphertext, sk SK, q int, p int64, mu int) []int64 {

	w := vecMod(vecSub(ct.C2, vecMatProd(ct.C1, sk.R1)), q)
	fmt.Println("w")
	fmt.Println(w)
	e := vecMod(w, int(p))
	fmt.Println("e")
	fmt.Println(e)

	v := vecMod(e, mu)
	m := vecDiv(vecSub(e, v), mu)
	// BLABLABLA
	return m
}
