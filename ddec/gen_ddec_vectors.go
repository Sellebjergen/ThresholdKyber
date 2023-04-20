package ddec

import (
	"crypto/rand"
	"fmt"
	"os"

	"ThresholdKyber.com/m/hybrid"
	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
	"golang.org/x/crypto/sha3"
)

func Generate_test_vec(paramsK int) {
	k := make([]byte, 32)
	rand.Read(k)
	pk, _ := hybrid.K_h1(paramsK)
	z := make([]byte, 32)
	//s := kyberk2so.IndcpaUnpackPrivateKey(sk, paramsK)

	coins := make([]byte, 32)
	ct, _ := kyberk2so.IndcpaEncrypt(k, pk, coins, paramsK)

	h := sha3.Sum256(pk)
	hash_c := sha3.Sum256(ct)

	u, v := kyberk2so.IndcpaUnpackCiphertext(ct, paramsK)
	t, seed := kyberk2so.IndcpaUnpackPublicKey(pk, paramsK)
	A, _ := kyberk2so.IndcpaGenMatrix(seed, false, paramsK)

	file_key_expected, err := os.Create("C:/Users/kaspe/Desktop/Speciale/ThresholdKyber/ddec/test_vectors_ddec/test_vector1/expected_output") // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer file_key_expected.Close()

	file_public_input, err := os.Create("C:/Users/kaspe/Desktop/Speciale/mp-spdz-0.3.5/Programs/Public-Input/kyber_ddec") // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer file_public_input.Close()

	/* file_private_input, err := os.Create("C:/Users/Kasper/Desktop/Speciale/mp-spdz-0.3.5/Player-Data/Input-P0-0") // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer file_private_input.Close() */

	//WritePolyVec(s, file_private_input)
	WriteExpectedKey(k, file_key_expected)
	WritePolyVec(u, file_public_input)
	WritePoly(v, file_public_input)
	for i := 0; i < paramsK; i++ {
		WritePolyVec(A[i], file_public_input)
	}
	WritePolyVec(t, file_public_input)
	WriteBytes(hash_c[:], file_public_input)

	file_s_expected, err := os.Create("C:/Users/kaspe/Desktop/Speciale/mp-spdz-0.3.5/Player-Data/Input-P0-0") // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
	}
	defer file_s_expected.Close()
	WriteBytes(z, file_s_expected)
	WriteBytes(h[:], file_s_expected)

}

func WritePolyVec(s kyberk2so.PolyVec, f *os.File) {
	for _, poly := range s {
		for i, coef := range poly {
			if i > 255 {
				break
			}
			f.WriteString(fmt.Sprintf("%d ", coef))
		}
		f.WriteString(fmt.Sprintf("\n"))
	}

}

func WritePoly(poly kyberk2so.Poly, f *os.File) {
	for i, coef := range poly {
		if i > 255 {
			break
		}
		f.WriteString(fmt.Sprintf("%d ", coef))
	}
	f.WriteString(fmt.Sprintf("\n"))
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

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
