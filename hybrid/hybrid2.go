package hybrid

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
	"golang.org/x/crypto/sha3"
)

type hybrid2Ciphertext struct {
	c1 []byte
	c2 []byte
	c3 []byte
	c4 []byte
}

func K_h2(paramsK int) ([]byte, []byte) {
	sk, pk, _ := kyberk2so.IndcpaKeypair(paramsK)
	return pk, sk
}

func E_h2(pk []byte, msg []byte, paramsK int) *hybrid2Ciphertext {
	k := make([]byte, 32)
	rand.Read(k)
	k_other_font := H(k)

	mu := H_prime(k)

	coins := make([]byte, 32)
	rand.Read(coins)
	c1, _ := kyberk2so.IndcpaEncrypt(k, pk, coins, paramsK)

	r := make([]byte, aes.BlockSize) // Randomness should be of length aes.BlockSize
	rand.Read(r)
	c2 := E_s(k_other_font, msg, r)

	c3 := G(make([]byte, 0), c2, mu)

	c4 := H_prime_prime(k)

	return &hybrid2Ciphertext{c1, c2, c3, c4}
}

func D_h2(sk []byte, ct *hybrid2Ciphertext, paramsK int) ([]byte, []byte) {
	k := kyberk2so.IndcpaDecrypt(ct.c1, sk, paramsK)
	if k == nil {
		return nil, nil
	}

	t := H_prime_prime(k)

	if !bytes.Equal(t, ct.c4) {
		return nil, nil
	}

	mu := H_prime(k)

	t_prime := G(make([]byte, 0), ct.c2, mu)

	if !bytes.Equal(t_prime, ct.c3) {
		return nil, nil
	}

	k_other_font := H(k)
	m := D_s(k_other_font, ct.c2)

	return k, m
}

func H_prime(in []byte) []byte {
	hash := sha3.New256()
	hash.Write(in)
	return hash.Sum(nil) // 32 bytes
}

func H_prime_prime(in []byte) []byte {
	hash := sha3.NewShake256()
	output := make([]byte, 32) // 32 is whatever key length we use for sym IND-CPA scheme
	hash.Write(in)
	hash.Read(output)
	return output
}
