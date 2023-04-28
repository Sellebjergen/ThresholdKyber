package gladius

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"reflect"
	"testing"
)

func TestGladiusEnc(t *testing.T) {
	params := InitParams(971, 2097143, 524288, 2, 512, 128)
	msg := make([]int64, params.N)
	msg[0] = 1
	msg[17] = 1

	pk, sk := keygen(params)
	ct := encrypt(params, msg, pk)
	res := decrypt(params, ct, sk)

	fmt.Println(res)
	fmt.Println(msg)

	if !reflect.DeepEqual(msg, res) {
		t.Errorf(("AAAAAAAAAAA"))
	}
}

func TestHybrid1Consistency(t *testing.T) {
	params := InitParams(971, 2097143, 524288, 2, 512, 128)
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	pk, sk := K_h1(params)
	ct := E_h1(params, pk, msg)
	_, msg_decrypted := D_h1(params, sk, ct)

	if !bytes.Equal(msg, msg_decrypted) {
		t.Errorf("Msg output by decrypt not equal to msg input to encrypt")
	}
}

func TestByteMsgConversion(t *testing.T) {
	params := InitParams(971, 2097143, 524288, 2, 512, 128)
	k := make([]byte, 32)
	rand.Read(k)

	interm := bytesToGladiusMsg(params, k)

	k_recomp := gladiusMsgToBytes(interm)

	fmt.Println(k)
	fmt.Println(interm)
	fmt.Println(k_recomp)

	if !reflect.DeepEqual(k, k_recomp) {
		t.Errorf(("AAAAAAAAAAA"))
	}
}
