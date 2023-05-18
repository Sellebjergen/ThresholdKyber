package main

import (
	"crypto/rand"
	"fmt"
	"reflect"
	"testing"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
	owcpa "ThresholdKyber.com/m/owcpa_TKyber"
)

func TestCPAEncryptDecrypt(t *testing.T) {
	msg := make([]byte, 32)
	rand.Read(msg)

	sk, pk, _ := kyberk2so.IndcpaKeypair(kyberk2so.ParamsK)

	coins := make([]byte, 32)
	rand.Read(coins)
	ct, _ := kyberk2so.IndcpaEncrypt(msg, pk, coins, kyberk2so.ParamsK)

	output_msg := kyberk2so.IndcpaDecrypt(ct, sk, kyberk2so.ParamsK)

	fmt.Println(output_msg)
	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Error: Decrypt(Encrypt(M)) != M")
	}
}

func TestEncDecNoCompression(t *testing.T) {
	msg := make([]byte, 32)
	rand.Read(msg)

	skpv, pkpv, Aseed := kyberk2so.IndcpaKeypair_nocomp(kyberk2so.ParamsK)

	coins := make([]byte, 32)
	rand.Read(coins)

	u, v := kyberk2so.IndcpaEncrypt_nocomp(msg, pkpv, Aseed, coins, kyberk2so.ParamsK)

	output_msg := kyberk2so.IndcpaDecrypt_nocomp(u, v, skpv, kyberk2so.ParamsK)
	fmt.Println(owcpa.Downscale(output_msg, 2, kyberk2so.ParamsQ))

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Error: Decrypt(Encrypt(M)) != M")
	}
}

func TestBarretReduceOtherModulus(t *testing.T) {
	res := kyberk2so.ByteopsBarrettReduce(int16(8000))
	fmt.Println(res)
	t.Errorf("AAAAAAAA")
}

func TestMontgReduceOtherModulus(t *testing.T) {
	res := kyberk2so.ByteopsMontgomeryReduce(int32(8000))
	fmt.Println(res)
	t.Errorf("AAAAAAAA")
}

func TestByteOpsCondSub(t *testing.T) {
	res := kyberk2so.ByteopsCSubQ(16000)
	fmt.Println(res)
	t.Errorf("AAAAAAAA")
}

func TestNTT(t *testing.T) {
	pv1 := kyberk2so.PolyvecNew(2)
	pv2 := kyberk2so.PolyvecNew(2)

	pv1[0][0] = 17
	pv1[0][1] = 42

	pv2[0][0] = 2
	pv2[0][1] = 1

	kyberk2so.PolyvecNtt(pv1, 2)
	kyberk2so.PolyvecNtt(pv2, 2)

	res := kyberk2so.PolyvecPointWiseAccMontgomery(pv1, pv2, 2)

	res = kyberk2so.PolyInvNttToMont(res)

	fmt.Println(res)
	t.Errorf("AAAAAAAA")
}

func TestInvNTTIsInverse(t *testing.T) {
	pv1 := kyberk2so.PolyvecNew(2)

	pv1[0][0] = 34
	pv1[0][1] = 101
	pv1[0][2] = 42

	kyberk2so.PolyvecNtt(pv1, 2)

	kyberk2so.PolyvecInvNttToMont(pv1, kyberk2so.ParamsK)

	for i, elem := range pv1[0] {
		pv1[0][i] = kyberk2so.Montgomery_reduce(int32(elem))
	}

	fmt.Println(pv1)
	t.Errorf("AAAAAAAA")
}

func TestOwnNTT(t *testing.T) {
	pv1 := new(Poly)

	fmt.Println(pv1)

	pv1[0] = 17
	pv1[1] = 101
	pv1[2] = 42

	pv1.nttGeneric()

	pv1.invNTTGeneric()

	fmt.Println(pv1)

	t.Errorf("AAAAAAAA")
}

func BenchmarkKyberKeygen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kyberk2so.IndcpaKeypair(kyberk2so.ParamsK)
	}
}

func BenchmarkKyberEncrypt(b *testing.B) {
	randMsg := make([]byte, 32)
	coins := make([]byte, 32)
	rand.Read(coins)
	rand.Read(randMsg)
	_, pk, _ := kyberk2so.IndcpaKeypair(kyberk2so.ParamsK)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		kyberk2so.IndcpaEncrypt(randMsg, pk, coins, kyberk2so.ParamsK)
	}
}

func BenchmarkKyberDecrypt(b *testing.B) {
	randMsg := make([]byte, 32)
	coins := make([]byte, 32)
	rand.Read(coins)
	rand.Read(randMsg)
	sk, pk, _ := kyberk2so.IndcpaKeypair(kyberk2so.ParamsK)
	ct, _ := kyberk2so.IndcpaEncrypt(randMsg, pk, coins, kyberk2so.ParamsK)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		kyberk2so.IndcpaDecrypt(ct, sk, kyberk2so.ParamsK)
	}
}
