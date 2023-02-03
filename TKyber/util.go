package TKyber

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
