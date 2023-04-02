package main

import (
	"fmt"
	"os"

	"ThresholdKyber.com/m/hybrid"
	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
)

func main() {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	pk, sk := hybrid.K_h1(2)

	s := kyberk2so.IndcpaUnpackPrivateKey(sk, 2)

	fmt.Println("s")
	fmt.Println(s)

	ct := hybrid.E_h1(pk, msg, 2, false)

	k := kyberk2so.IndcpaDecrypt(ct.C1, sk, 2)

	u, v := kyberk2so.IndcpaUnpackCiphertext(ct.C1, 2)

	fmt.Println("Generating test vector for Kyber DDec")
	fmt.Println("k")
	fmt.Println(k)
	fmt.Println("kyber_ct_1")
	fmt.Println(u)
	fmt.Println("kyber_ct_2")
	fmt.Println(v)
	fmt.Println("c2")
	fmt.Println(ct.C2)
	fmt.Println("c3")
	fmt.Println(ct.C3)

	write_ddec_input(s, u, v, ct.C2, ct.C3)
}

func write_ddec_input(s kyberk2so.PolyVec, u kyberk2so.PolyVec, v kyberk2so.Poly, c2 []byte, c3 []byte) {
	file_public_input, err := os.Create("C:/Users/Kasper/Desktop/Speciale/mp-spdz-0.3.5/Programs/Public-Input/kyber_ddec") // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer file_public_input.Close()

	file_private_input, err := os.Create("C:/Users/Kasper/Desktop/Speciale/mp-spdz-0.3.5/Player-Data/Input-P0-0") // creating...
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer file_private_input.Close()

	writePolyVec(s, file_private_input)
	writePolyVec(u, file_public_input)
	writePoly(v, file_public_input)
	writeBytes(c2, file_public_input)
	writeBytes(c3, file_public_input)
}

func writePolyVec(s kyberk2so.PolyVec, f *os.File) {
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

func writePoly(poly kyberk2so.Poly, f *os.File) {
	for i, coef := range poly {
		if i > 255 {
			break
		}
		f.WriteString(fmt.Sprintf("%d ", coef))
	}
	f.WriteString(fmt.Sprintf("\n"))
}

func writeBytes(b []byte, f *os.File) {
	for _, oneByte := range b {
		byte_as_string := fmt.Sprintf("%08b", oneByte)
		to_write := ""
		for _, char := range byte_as_string {
			to_write += string(char) + " "
		}
		reverse(to_write)
		f.WriteString(to_write)
	}
	f.WriteString(fmt.Sprintf("\n"))
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
