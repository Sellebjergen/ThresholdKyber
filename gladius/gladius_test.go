package gladius

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGladiusEnc(t *testing.T) {

	n := 971
	msg := make([]int64, n)
	msg[0] = 1
	msg[17] = 1
	q := 2097143
	ell := int64(524288)
	t_param := 2
	p := int64(512)
	mu := 128

	pk, sk := keygen(n, q, ell)
	ct := encrypt(n, q, ell, t_param, p, msg, pk)

	fmt.Println(ct)

	res := decrypt(ct, sk, q, p, mu)

	fmt.Println(res)

	if !reflect.DeepEqual(msg, res) {
		t.Errorf(("AAAAAAAAAAA"))
	}
}
