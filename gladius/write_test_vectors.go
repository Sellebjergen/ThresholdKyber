package gladius

import (
	"fmt"
	"os"
)

func WriteDDecTestVec() {

	n := 100
	msg := make([]int64, n)
	msg[0] = 1
	msg[17] = 1
	q := 2097143
	ell := int64(524288)
	t_param := 2
	p := int64(512)
	//mu := 128

	pk, sk := keygen(n, q, ell)
	ct := encrypt(n, q, ell, t_param, p, msg, pk)

	file_private_input, err := os.Create("C:/Users/kasper/Desktop/Speciale/mp-spdz-0.3.5/Player-Data/Input-P0-0") // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
	}
	defer file_private_input.Close()

	file_public_input, err := os.Create("C:/Users/kasper/Desktop/Speciale/mp-spdz-0.3.5/Programs/Public-Input/gladius_ddec") // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
	}
	defer file_public_input.Close()

	WriteVec(ct.C1, file_public_input)
	WriteVec(ct.C2, file_public_input)

	WriteMat(sk.R1, file_private_input)
}

func WriteMat(mat [][]int64, f *os.File) {
	for i := range mat {
		for _, coef := range mat[i] {
			f.WriteString(fmt.Sprintf("%d ", coef))
		}
	}
	f.WriteString(fmt.Sprintf("\n"))
}

func WriteVec(vec []int64, f *os.File) {
	for _, coef := range vec {
		f.WriteString(fmt.Sprintf("%d ", coef))
	}
	f.WriteString(fmt.Sprintf("\n"))
}
