package gladius

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"

	"golang.org/x/crypto/sha3"
)

type Hybrid1Ciphertext struct {
	C1 Ciphertext
	C2 []byte
	C3 []byte
}

func K_h1(params *GladiusParams) (PK, SK) {
	pk, sk := keygen(params)
	return pk, sk
}

func E_h1(params *GladiusParams, pk PK, msg []byte) *Hybrid1Ciphertext {
	k := make([]byte, 32)
	rand.Read(k)
	k_other_font := H(k)

	c1 := encrypt(params, bytesToGladiusMsg(params, k), pk)

	r := make([]byte, aes.BlockSize) // Randomness should be of length aes.BlockSize
	rand.Read(r)
	c2 := E_s(k_other_font, msg, r)

	c3 := G(make([]byte, 0), c2, k)

	return &Hybrid1Ciphertext{c1, c2, c3}
}

func D_h1(params *GladiusParams, sk SK, ct *Hybrid1Ciphertext) ([]byte, []byte) {
	k := decrypt(params, ct.C1, sk)
	if k == nil {
		return nil, nil
	}

	k_bytes := gladiusMsgToBytes(k)

	t := G(make([]byte, 0), ct.C2, k_bytes)

	if !bytes.Equal(t, ct.C3) {
		return nil, nil
	}

	k_other_font := H(k_bytes)
	m := D_s(k_other_font, ct.C2)

	return k_bytes, m
}

func H(in []byte) []byte {
	hash := sha3.NewShake256()
	output := make([]byte, 32) // 32 is whatever key length we use for sym IND-CPA scheme
	hash.Write(in)
	hash.Read(output)
	return output
}

func G(ct []byte, inBytes []byte, inMsg []byte) []byte {
	hash := sha3.New256()
	hash.Write(ct)
	hash.Write(inBytes)
	hash.Write(inMsg)
	return hash.Sum(nil)
}
