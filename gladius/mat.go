package gladius

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
	randmath "math/rand"

	"ThresholdKyber.com/m/util"
)

type Mat [][]int64

func sampleRMatrix(n int) Mat {
	mat := make(Mat, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int64, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ai := randmath.Intn(2)
			bi := randmath.Intn(2)
			mat[i][j] = int64(ai - bi)
		}
	}
	return mat
}

func sampleUniform(n int, q int) Mat {
	mat := make([][]int64, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int64, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			val, _ := rand.Int(rand.Reader, big.NewInt(int64(q)))
			mat[i][j] = int64(val.Int64())
		}
	}
	return mat
}

func constructGadgetMat(ell int64, n int) Mat {
	mat := make([][]int64, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int64, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				mat[i][j] = ell
			}

		}
	}
	return mat
}

func matAdd(a Mat, b Mat) Mat {
	res := make([][]int64, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = make([]int64, len(a[0]))
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			res[i][j] = a[i][j] + b[i][j]
		}
	}
	return res
}

func vecMatProd(a []int64, b Mat) []int64 {
	n := len(a)
	res := make([]int64, n)
	for j := 0; j < n; j++ {
		sum := int64(0)
		for k := 0; k < n; k++ {
			sum = sum + a[k]*b[k][j]
		}
		res[j] = sum
	}
	return res
}

func vecMod(a []int64, modulus int) []int64 {
	res := make([]int64, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = int64(util.Euc_mod(int(a[i]), modulus))
		if res[i] > int64(modulus)/2 {
			res[i] = res[i] - int64(modulus) // Centre
		}
	}
	return res
}

func vecSub(a []int64, b []int64) []int64 {
	res := make([]int64, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[i] - b[i]
	}
	return res
}

func vecDiv(a []int64, factor int) []int64 {
	res := make([]int64, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[i] / int64(factor)
	}
	return res
}

func matProd(a Mat, b Mat, n int) Mat {
	res := make([][]int64, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int64, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				res[i][j] = res[i][j] + a[i][k]*b[k][j]

			}
		}
	}

	return res
}

func round_mod(x int64, p int64, q int) int64 {
	scale := float64(x) * float64(p) / float64(q)
	rounded := math.Round(scale)
	fmt.Println(rounded)
	mod := int64(util.Euc_mod(int(rounded), int(p)))
	if mod > int64(q)/2 {
		return mod - int64(q) // Centre
	}
	return mod
}

func vec_round_mod(a []int64, p int64, q int) []int64 {
	res := make([]int64, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = round_mod(a[i], p, q)
	}
	return res
}
