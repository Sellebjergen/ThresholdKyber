package util

import "ThresholdKyber.com/m/kyber"

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

func Transpose(slice [][]*kyber.Poly) [][]*kyber.Poly {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]*kyber.Poly, xl)
	for i := range result {
		result[i] = make([]*kyber.Poly, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}
