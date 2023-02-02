package TKyber

// Regular max function for ints
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Euclidean modulo function
func euc_mod(x, m int32) int32 {
	res := x % m
	if res < 0 {
		res += m
	}
	return res
}

// Reverses slice of int32s
func Reverse(input []int32) []int32 {
	var output []int32

	for i := len(input) - 1; i >= 0; i-- {
		output = append(output, input[i])
	}

	return output
}
