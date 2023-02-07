package indcpa_TKyber

import (
	"ThresholdKyber.com/m/kyber"
	owcpa "ThresholdKyber.com/m/owcpa_TKyber"
)

func Setup(params kyber.ParameterSet, n int, t int) (*kyber.IndcpaPublicKey, []kyber.PolyVec) {
	return owcpa.Setup(params, n, t)
}

func PartDec() {

}

func Combine() {

}

func Enc() {

}
