package hybrid

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
	"golang.org/x/crypto/sha3"
)

type Hybrid1Ciphertext struct {
	C1 []byte
	C2 []byte
	C3 []byte
}

func K_h1(paramsK int) ([]byte, []byte) {
	sk, pk, _ := kyberk2so.IndcpaKeypair(paramsK)
	return pk, sk
}

func E_h1(pk []byte, msg []byte, paramsK int, isDet bool) *Hybrid1Ciphertext {
	k := make([]byte, 32)
	rand.Read(k)
	k_other_font := H(k)

	coins := make([]byte, 32)
	rand.Read(coins)
	c1, _ := kyberk2so.IndcpaEncrypt(k, pk, coins, paramsK)

	r := make([]byte, aes.BlockSize) // Randomness should be of length aes.BlockSize
	rand.Read(r)
	c2 := E_s(k_other_font, msg, r)

	var c3 []byte
	if isDet {
		c3 = G(make([]byte, 0), c2, k)
	} else {
		c3 = G(c1, c2, k)
	}

	return &Hybrid1Ciphertext{c1, c2, c3}
}

func D_h1(sk []byte, ct *Hybrid1Ciphertext, paramsK int, isDet bool) ([]byte, []byte) {
	k := kyberk2so.IndcpaDecrypt(ct.C1, sk, paramsK)
	if k == nil {
		return nil, nil
	}

	var t []byte
	if isDet {
		t = G(make([]byte, 0), ct.C2, k)
	} else {
		t = G(ct.C1, ct.C2, k)
	}

	if !bytes.Equal(t, ct.C3) {
		return nil, nil
	}

	k_other_font := H(k)
	m := D_s(k_other_font, ct.C2)

	return k, m
}

func H(in []byte) []byte {
	hash := sha3.NewShake256()
	output := make([]byte, 32) // 32 is whatever key length we use for sym IND-CPA scheme
	hash.Write(in)
	hash.Read(output)
	return output
}

func G(ct []byte, inBytes []byte, inMsg []byte) []byte {
	hash := sha3.NewShake256()
	output := make([]byte, 256) // 256 is placeholder for |G|
	hash.Write(ct)
	hash.Write(inBytes)
	hash.Write(inMsg)
	hash.Read(output)
	return output
}
