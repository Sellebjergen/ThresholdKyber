package util

import (
	"fmt"

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
		fmt.Println(newRes)
		fmt.Println(newTmp)
		return newRes, newTmp
	}

	for i := left; i < n+1; i++ {
		newTmp = append(newTmp, i)
		newRes, newTmp = makeCombRecursive(n, i+1, t-1, newRes, newTmp)

		newTmp = newTmp[:len(newTmp)-1]
	}

	return newRes, newTmp
}
