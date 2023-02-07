package TKyber

import (
	"fmt"
	"reflect"
	"testing"

	"ThresholdKyber.com/m/kyber"
)

func TestSetupWorksInCaseNis3(t *testing.T) {
	// TODO: sk_shares has a wrong coefficient on the very first index - the rest of the polynomial seem correct
	pk, sk_shares := Setup(*kyber.Kyber512, 3, 3)

	// total of first share
	var sk1 kyber.Poly
	sk1.Add(&sk1, sk_shares[0].Vec[0])
	sk1.Add(&sk1, sk_shares[1].Vec[0])
	sk1.Add(&sk1, sk_shares[2].Vec[0])

	// total of second share
	var sk2 kyber.Poly
	sk2.Add(&sk2, sk_shares[0].Vec[1])
	sk2.Add(&sk2, sk_shares[1].Vec[1])
	sk2.Add(&sk2, sk_shares[2].Vec[1])

	// assemble secret key
	sk := kyber.PolyVec{Vec: []*kyber.Poly{&sk1, &sk2}}
	r := make([]byte, kyber.Kyber512.IndcpaSecretKeySize)
	kyber.PackSecretKey(r, &sk)
	skPacked := &kyber.IndcpaSecretKey{
		Packed: r,
	}

	// calling the encrypt decrypt functionality, to check that the key works as expected.
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	out := make([]byte, 32)
	ct := make([]byte, kyber.Kyber512.CipherTextSize())
	coins := make([]byte, 32)
	kyber.Kyber512.IndcpaEncrypt(ct, msg, pk, coins)
	kyber.Kyber512.IndcpaDecrypt(out, ct, skPacked)

	if !reflect.DeepEqual(msg, out) {
		t.Errorf("Decryption failed!")
	}
}

// ================= Integration tests =================

func TestSimpleCase(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	pk, sk_shares := Setup(*kyber.Kyber512, 3, 3)

	coins := make([]byte, 32)
	// rand.Read(coins)

	ct := make([]byte, kyber.Kyber512.CipherTextSize())
	kyber.Kyber512.IndcpaEncrypt(ct, msg, pk, coins)

	// Decrypt
	d_1 := PartDec(*kyber.Kyber512, sk_shares[0], ct, 0)
	d_2 := PartDec(*kyber.Kyber512, sk_shares[1], ct, 1)
	d_3 := PartDec(*kyber.Kyber512, sk_shares[2], ct, 2)

	fmt.Println(d_1)

	combined := Combine(ct, d_1, d_2, d_3)
	fmt.Println(combined)

	output_msg := make([]byte, 32)
	combined.ToMsg(output_msg)

	fmt.Println(msg)
	fmt.Println(output_msg)

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Error")
	}
}

/* func TestSetupUsing1PlayerGivesBackSecretKey(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	pk, skShares := Setup(*kyber.Kyber512, 1, 1)
	polyVec := kyber.Kyber512.AllocPolyVec()

	m := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	coins := make([]byte, 32)
	c := make([]byte, kyber.Kyber512.CipherTextSize())
	kyber.Kyber512.IndcpaEncrypt(c, m, pk, coins)

	out := make([]byte, 32)
	r := make([]byte, kyber.Kyber512.IndcpaSecretKeySize)
	kyber.PackSecretKey(r, &polyVec)
	sk := kyber.IndcpaSecretKey{Packed: r}
	kyber.Kyber512.IndcpaDecrypt(out, c, &sk)

	if !reflect.DeepEqual(msg, out) {
		t.Errorf("Error")
	}
} */

func TestFullWithN1(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	pk, skShares := Setup(*kyber.Kyber512, 1, 1)

	coins := make([]byte, 32)

	ct := make([]byte, kyber.Kyber512.CipherTextSize())
	kyber.Kyber512.IndcpaEncrypt(ct, msg, pk, coins)

	d1 := PartDec(*kyber.Kyber512, skShares[0], ct, 0)
	combined := Combine(ct, d1)

	output_msg := make([]byte, 32)
	combined.ToMsg(output_msg)

	fmt.Println(combined)

	fmt.Println(msg)
	fmt.Println(output_msg)

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Error")
	}
}
