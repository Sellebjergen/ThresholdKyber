package gladius

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

type GladiusParams struct {
	N   int
	Q   int
	Ell int64
	T   int
	P   int64
	Mu  int
}

func InitParams(n int, q int, ell int64, t int, p int64, mu int) *GladiusParams {
	return &GladiusParams{n, q, ell, t, p, mu}
}

func keygen(params *GladiusParams) (PK, SK) {
	R1 := sampleRMatrix(params.N)
	R2 := sampleRMatrix(params.N)
	A1 := sampleUniform(params.N, params.Q)
	G := constructGadgetMat(params.Ell, params.N)

	A2 := matAdd(matAdd(matProd(A1, R1, params.N), R2), G)

	pk := PK{A1, A2}
	sk := SK{R1, pk}

	return pk, sk

}

func encrypt(params *GladiusParams, msg []int64, pk PK) Ciphertext {
	c1 := vec_round_mod(vecMatProd(msg, pk.A1), params.P, params.Q)
	c2 := vec_round_mod(vecMatProd(msg, pk.A2), params.P, params.Q)
	return Ciphertext{c1, c2}
}

func decrypt(params *GladiusParams, ct Ciphertext, sk SK) []int64 {

	w := vecMod(vecSub(ct.C2, vecMatProd(ct.C1, sk.R1)), params.Q)
	e := vecMod(w, int(params.P))

	v := vecMod(e, params.Mu)
	m := vecDiv(vecSub(e, v), params.Mu)
	// BLABLABLA
	return m
}
