package main

import (
	"crypto/rand"
	"reflect"
	"testing"

	"ThresholdKyber.com/m/kyber"
)

func TestCPAEncryptDecrypt(t *testing.T) {
	ct := make([]byte, kyber.Kyber512.CipherTextSize())
	msg := make([]byte, 32)
	rand.Read(msg)
	pk, sk, _ := kyber.Kyber512.IndcpaKeyPair(rand.Reader)

	coins := make([]byte, 32)
	rand.Read(coins)

	kyber.Kyber512.IndcpaEncrypt(ct, msg, pk, coins)

	// Decrypt
	output_msg := make([]byte, 32)
	kyber.Kyber512.IndcpaDecrypt(output_msg, ct, sk)

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Error: Decrypt(Encrypt(M)) != M")
	}
}
