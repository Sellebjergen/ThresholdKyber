package gladius

import (
	"crypto/aes"
	"crypto/cipher"
)

func E_s(key []byte, msg []byte, randomness []byte) []byte {
	if len(randomness) != aes.BlockSize {
		panic("Randomness not of the correct length")
	}
	block, aesErr := aes.NewCipher(key)
	if aesErr != nil {
		panic(aesErr)
	}

	ct := make([]byte, aes.BlockSize+len(msg))
	streamCTR := cipher.NewCTR(block, randomness)

	streamCTR.XORKeyStream(ct[aes.BlockSize:], msg)
	copy(ct[:aes.BlockSize], randomness)
	return ct
}

func D_s(key []byte, ct []byte) []byte {
	block, aesErr := aes.NewCipher(key)
	if aesErr != nil {
		panic(aesErr)
	}

	IV := ct[:aes.BlockSize]
	streamCTR := cipher.NewCTR(block, IV)
	plaintext := make([]byte, len(ct)-aes.BlockSize)

	streamCTR.XORKeyStream(plaintext, ct[aes.BlockSize:])
	return plaintext
}
