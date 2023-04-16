package large_mod

import (
	"fmt"
	"math"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
	"ThresholdKyber.com/m/util"
)

func Merge(poly_1 kyberk2so.Poly, poly_2 kyberk2so.Poly, q_1 int, q_2 int) []int {
	N := q_1 * q_2

	// Euclid's Extended Algorithm
	res := make([]int, kyberk2so.ParamsN)
	for i := 0; i < kyberk2so.ParamsN; i++ {
		fmt.Println("Input")
		fmt.Println(poly_1[i])
		fmt.Println(poly_2[i])
		_, x, y := euclidsExtendedAlgorithm(float64(q_1), float64(q_2))

		fmt.Println("x")
		fmt.Println(x)
		fmt.Println("y")
		fmt.Println(y)

		// Solve CRT using res of EEA
		res[i] = util.Euc_mod(q_1*x*int(poly_2[i])+q_2*y*int(poly_1[i]), N)

		fmt.Println("Res")
		fmt.Println(res[i])
	}

	fmt.Println(res)

	// Scaling and rounding
	for i := 0; i < kyberk2so.ParamsN; i++ {
		res[i] = int(math.Round((float64(2) / float64(N)) * float64(res[i])))
	}
	for i := 0; i < kyberk2so.ParamsN; i++ {
		res[i] = res[i] % 2
	}

	return res
}

func euclidsExtendedAlgorithm(a float64, b float64) (int, int, int) {
	d_0 := a
	x_0 := float64(1)
	y_0 := float64(0)

	d_1 := b
	x_1 := float64(0)
	y_1 := float64(1)
	for d_1 != 0 {
		q := math.Floor(d_0 / d_1)
		d_2 := d_1
		x_2 := x_1
		y_2 := y_1

		d_1 = d_0 - q*d_1
		x_1 = x_0 - q*x_1
		y_1 = y_0 - q*y_1

		d_0 = d_2
		x_0 = x_2
		y_0 = y_2
	}

	return int(d_0), int(x_0), int(y_0) // not d but either x or y
}
