package gladius

import (
	"fmt"
	"os"
)

func WriteDDecTestVec() {
	params := InitParams(256, 2097143, 524288, 2, 512, 128)

	msg := make([]int64, params.N)
	msg[0] = 1
	msg[17] = 1

	msg_bytes := gladiusMsgToBytes(msg)

	pk, sk := K_h1(params)
	ct := E_h1(params, pk, msg_bytes)
	k, msg_decrypted := D_h1(params, sk, ct)

	fmt.Println(k)
	fmt.Println(msg_decrypted)

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

	file_key, err := os.Create("C:/Users/kasper/Desktop/expected_key") // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
	}
	defer file_key.Close()

	WriteExpectedKey(k, file_key)

	WriteVec(ct.C1.C1, file_public_input)
	WriteVec(ct.C1.C2, file_public_input)
	WriteMat(pk.A1, file_public_input)
	WriteMat(pk.A2, file_public_input)
	WriteBytes(ct.C2, file_public_input)
	WriteBytes(ct.C3, file_public_input)

	WriteMat(sk.R1, file_private_input)
}

func WriteBytes(b []byte, f *os.File) {
	for _, oneByte := range b {
		byte_as_string := fmt.Sprintf("%08b", oneByte)
		to_write := ""
		for _, char := range reverse(byte_as_string) {
			to_write += string(char) + " "
		}
		f.WriteString(to_write)
	}
	f.WriteString(fmt.Sprintf("\n"))
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

func WriteExpectedKey(b []byte, f *os.File) {
	f.WriteString(fmt.Sprintf("["))
	for i, oneByte := range b {
		byte_as_string := fmt.Sprintf("%08b", oneByte)
		to_write := ""
		for j, char := range reverse(byte_as_string) {
			if (j != 7) || i != len(b)-1 {
				to_write += string(char) + ", "
			} else {
				to_write += string(char)
			}

		}
		f.WriteString(to_write)
	}
	f.WriteString(fmt.Sprintf("]"))
}
