package main

import (
	"crypto/rand"
	"reflect"
	"testing"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
)

func TestCPAEncryptDecrypt(t *testing.T) {
	msg := make([]byte, 32)
	rand.Read(msg)

	sk, pk, _ := kyberk2so.IndcpaKeypair(kyberk2so.ParamsK)

	coins := make([]byte, 32)
	rand.Read(coins)
	ct, _ := kyberk2so.IndcpaEncrypt(msg, pk, coins, kyberk2so.ParamsK)

	output_msg := kyberk2so.IndcpaDecrypt(ct, sk, kyberk2so.ParamsK)

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Error: Decrypt(Encrypt(M)) != M")
	}
}
