package gladius

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"testing"
)

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
