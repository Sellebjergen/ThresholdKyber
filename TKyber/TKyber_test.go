package TKyber

import (
	"crypto/rand"
	"fmt"
	"reflect"
	"testing"

	"ThresholdKyber.com/m/kyber"
)

// ================= Integration tests =================

func TestSimpleCase(t *testing.T) {
	rq := new(quotRing).initKyberRing()
	ct := make([]byte, kyber.Kyber512.CipherTextSize())
	msg := make([]byte, 32)
	rand.Read(msg)
	pk, sk_shares := Setup(*kyber.Kyber512, 3, 3)

	coins := make([]byte, 32)
	rand.Read(coins)

	kyber.Kyber512.IndcpaEncrypt(ct, msg, pk, coins)

	// Decrypt
	d_1 := rq.PartDec(*kyber.Kyber512, sk_shares[0], ct, 0)
	d_2 := rq.PartDec(*kyber.Kyber512, sk_shares[1], ct, 1)
	d_3 := rq.PartDec(*kyber.Kyber512, sk_shares[2], ct, 2)

	combined := rq.Combine(ct, d_1, d_2, d_3)
	fmt.Println(combined)

	output_msg := make([]byte, 32)

	combined.ToMsg(output_msg)

	fmt.Println(msg)
	fmt.Println(output_msg)

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Error")
	}
}
