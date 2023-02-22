package hybrid

import (
	"bytes"
	"testing"
)

func TestHybrid2Consistency(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	pk, sk := K_h2(2)

	ct := E_h2(pk, msg, 2)

	_, msg_decrypted := D_h2(sk, ct, 2)

	if !bytes.Equal(msg, msg_decrypted) {
		t.Errorf("Msg output by decrypt not equal to msg input to encrypt")
	}
}
