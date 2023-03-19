package main

import (
	"fmt"

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
}
