package TKyber

import (
	"ThresholdKyber.com/m/kyber"
)

type SKshare struct {
}

func Share(params kyber.ParameterSet, sk *kyber.IndcpaSecretKey) []*SKshare {
	polyVec_sk := params.AllocPolyVec()
	kyber.UnpackSecretKey(&polyVec_sk, sk.Packed)

	return nil
}

func Rec() {

}
