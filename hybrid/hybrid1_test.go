package hybrid

import (
	"bytes"
	"testing"
)

func TestHybrid1Consistency(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	pk, sk := K_h1(2)

	coins := make([]byte, 32)
	ct := E_h1(pk, msg, 2, coins, false)

	_, msg_decrypted := D_h1(sk, ct, 2, false)

	if !bytes.Equal(msg, msg_decrypted) {
		t.Errorf("Msg output by decrypt not equal to msg input to encrypt")
	}
}
