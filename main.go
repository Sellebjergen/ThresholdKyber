package main

import (
	"ThresholdKyber.com/m/ddec"
	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
)

func main() {
	//ddec.Generate_test_vec(2)
	ddec.Generate_Enc_Vec(kyberk2so.ParamsK)
	//gladius.WriteDDecTestVec()
}
