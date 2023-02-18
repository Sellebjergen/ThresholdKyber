package util

import (
	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
)

// Regular max function for ints
func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

// Euclidean modulo function
func euc_mod(x, m int) int {
	res := x % m

	if res < 0 {
		res += m
	}

	return res
}

// Reverses slice of ints
func Reverse(input []int) []int {
	var output []int

	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}

	return output
}

func Transpose(slice [][]kyberk2so.Poly) [][]kyberk2so.Poly {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]kyberk2so.Poly, xl)
	for i := range result {
		result[i] = make([]kyberk2so.Poly, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func SwapFirstAndSecondDim(slice [][][]kyberk2so.Poly) [][][]kyberk2so.Poly {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][][]kyberk2so.Poly, xl)
	for i := range result {
		result[i] = make([][]kyberk2so.Poly, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func MakeCombinations(n int, t int) [][]int {
	res := make([][]int, 0)
	tmp := make([]int, 0)

	res, _ = makeCombRecursive(n, 1, t, res, tmp)
	return res
}

func makeCombRecursive(n int, left int, t int, res [][]int, tmp []int) ([][]int, []int) {
	newRes := make([][]int, len(res))
	newTmp := make([]int, len(tmp))
	copy(newRes, res)
	copy(newTmp, tmp)
	if t == 0 {
		toAdd := make([]int, len(newTmp))
		copy(toAdd, newTmp)
		newRes = append(newRes, toAdd)

		return newRes, newTmp
	}

	for i := left; i < n+1; i++ {
		newTmp = append(newTmp, i)
		newRes, newTmp = makeCombRecursive(n, i+1, t-1, newRes, newTmp)

		newTmp = newTmp[:len(newTmp)-1]
	}

	return newRes, newTmp
}

func Contains(list []int, elem int) bool {
	for _, el := range list {
		if el == elem {
			return true
		}
	}
	return false
}

type Bit uint8

// Inefficient
func BytesToBits(b []byte) []Bit {
	result := make([]Bit, len(b)*8)
	for i := 0; i < len(b); i++ {
		result[i*8] = Bit(b[i] ^ 1)
		result[i*8+1] = Bit(b[i] ^ 2)
		result[i*8+2] = Bit(b[i] ^ 4)
		result[i*8+3] = Bit(b[i] ^ 8)
		result[i*8+4] = Bit(b[i] ^ 16)
		result[i*8+5] = Bit(b[i] ^ 32)
		result[i*8+6] = Bit(b[i] ^ 64)
		result[i*8+7] = Bit(b[i] ^ 128)
	}
	return result
}
