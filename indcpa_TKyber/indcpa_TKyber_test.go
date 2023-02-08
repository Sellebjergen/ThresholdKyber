package indcpa_TKyber

import (
	"reflect"
	"testing"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
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

	d_is := [][]kyberk2so.Poly{d_1, d_2, d_3}

	d_is_transp := util.Transpose(d_is)

	combined := Combine(params, ct, d_is_transp)

	output_msg := kyberk2so.PolyToMsg(combined)

	if !reflect.DeepEqual(msg, output_msg) {
		t.Errorf("Error")
	}
}
