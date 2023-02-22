package hybrid

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"testing"
)

func TestHybridConsistency(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	pk, sk := K_h1(2)

	ct := E_h1(pk, msg, 2, false)

	_, msg_decrypted := D_h1(sk, ct, 2, false)

	if !bytes.Equal(msg, msg_decrypted) {
		t.Errorf("Msg output by decrypt not equal to msg input to encrypt")
	}
}

func TestSymConsistency(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	sym_key := make([]byte, 32)
	rand.Read(sym_key)

	r := make([]byte, aes.BlockSize)
	rand.Read(r)

	ct := E_s(sym_key, msg, r)

	msg_decrypted := D_s(sym_key, ct)

	if !bytes.Equal(msg, msg_decrypted) {
		t.Errorf("Msg output by decrypt not equal to msg input to encrypt")
	}
}
