package TKyber

import (
	"fmt"
	"reflect"
	"testing"

	"ThresholdKyber.com/m/kyber"
)

func TestSetupWorksInCaseNis3(t *testing.T) {
	// TODO: sk_shares has a wrong coefficient on the very first index - the rest of the polynomial seem correct
	rq := new(quotRing).initKyberRing()
	pk, sk_shares := Setup(*kyber.Kyber512, 3, 3)

	// total of first share
	sk1 := &Polynomial{Coeffs: []int{0}}
	sk1 = rq.add(sk1, sk_shares[0][0])
	sk1 = rq.add(sk1, sk_shares[1][0])
	sk1 = rq.add(sk1, sk_shares[2][0])

	// total of second share
	sk2 := &Polynomial{Coeffs: []int{0}}
	sk2 = rq.add(sk2, sk_shares[0][1])
	sk2 = rq.add(sk2, sk_shares[1][1])
	sk2 = rq.add(sk2, sk_shares[2][1])

	// assemble secret key
	sk := kyber.PolyVec{Vec: []*kyber.Poly{sk1.toKyberPoly(), sk2.toKyberPoly()}}
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
	rq := new(quotRing).initKyberRing()
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	pk, sk_shares := Setup(*kyber.Kyber512, 3, 3)

	coins := make([]byte, 32)
	// rand.Read(coins)

	ct := make([]byte, kyber.Kyber512.CipherTextSize())
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

func TestFullWithN1(t *testing.T) {
	rq := new(quotRing).initKyberRing()
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	pk, sk_shares := Setup(*kyber.Kyber512, 1, 1)

	coins := make([]byte, 32)
	// rand.Read(coins)

	ct := make([]byte, kyber.Kyber512.CipherTextSize())
	kyber.Kyber512.IndcpaEncrypt(ct, msg, pk, coins)

	// Decrypt
	d_1 := rq.PartDec(*kyber.Kyber512, sk_shares[0], ct, 0)

	combined := rq.Combine(ct, d_1)
	fmt.Println(combined)

	output_msg := make([]byte, 32)
	combined.ToMsg(output_msg)

	fmt.Println(msg)
	fmt.Println(output_msg)

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Error")
	}
}
