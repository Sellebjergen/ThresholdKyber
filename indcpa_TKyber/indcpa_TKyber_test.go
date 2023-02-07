package indcpa_TKyber

import (
	"fmt"
	"reflect"
	"testing"

	"ThresholdKyber.com/m/kyber"
	owcpa "ThresholdKyber.com/m/owcpa_TKyber"
	"ThresholdKyber.com/m/util"
)

// ================= Integration tests =================

func TestSimpleCase(t *testing.T) {
	msg := []byte("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	params := owcpa.NewParameterSet("TKyber-Test")
	n := 3
	t_param := 3
	delta := 10
	pk, sk_shares := Setup(params, n, t_param)

	ct := Enc(params, msg, pk, delta)

	// Decrypt
	d_1 := PartDec(params, sk_shares[0], ct, 0, delta)
	d_2 := PartDec(params, sk_shares[1], ct, 1, delta)
	d_3 := PartDec(params, sk_shares[2], ct, 2, delta)

	d_is := [][]*kyber.Poly{d_1, d_2, d_3}

	d_is_transp := util.Transpose(d_is)

	fmt.Println(d_is_transp)

	combined := Combine(params, ct, d_is_transp)

	output_msg := make([]byte, 32)
	combined.ToMsg(output_msg)

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Error")
	}
}
