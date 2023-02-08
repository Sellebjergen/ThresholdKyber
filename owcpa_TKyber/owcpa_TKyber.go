package owcpa_TKyber

import (
	"crypto/rand"
	"math"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
)

func Setup(params *OwcpaParams, n int, t int) ([]byte, []kyberk2so.PolyVec) {
	// Run setup to get Kyber KeyPair
	sk, pk, _ := kyberk2so.IndcpaKeypair(kyberk2so.ParamsK)
	sk_unpacked := kyberk2so.IndcpaUnpackPrivateKey(sk, kyberk2so.ParamsK)

	// Perform secret sharing
	sk_shares := params.LSS_scheme.Share(sk_unpacked, n)

	return pk, sk_shares
}

func Enc(params *OwcpaParams, msg []byte, pk []byte) []byte {
	coins := make([]byte, 32)
	rand.Read(coins)
	ct, _ := kyberk2so.IndcpaEncrypt(msg, pk, coins, kyberk2so.ParamsK)
	return ct
}

func PartDec(params *OwcpaParams, sk_i kyberk2so.PolyVec, ct []byte, party int) kyberk2so.Poly {
	var zero kyberk2so.Poly
	// Sample noise
	e_i := params.D_flood_dist.SampleNoise(params.Q, 255, params.Sigma) // TODO: Fix params

	bp, v := kyberk2so.IndcpaUnpackCiphertext(ct, kyberk2so.ParamsK)

	// Inner prod
	kyberk2so.PolyvecNtt(bp, kyberk2so.ParamsK)
	d_i := kyberk2so.PolyvecPointWiseAccMontgomery(sk_i, bp, kyberk2so.ParamsK)

	d_i = kyberk2so.PolyInvNttToMont(d_i)
	if party == 0 {
		d_i = kyberk2so.PolySub(v, d_i)
	} else {
		d_i = kyberk2so.PolySub(zero, d_i)
	}

	// Add noise
	d_i = kyberk2so.PolyAdd(d_i, e_i)

	d_i = kyberk2so.PolyReduce(d_i)

	return d_i
}

func Combine(params *OwcpaParams, ct []byte, d_is []kyberk2so.Poly) kyberk2so.Poly {
	p := 2

	y := params.LSS_scheme.Rec(d_is)
	unrounded := make([]float64, len(y))

	for i := 0; i < len(unrounded); i++ {
		unrounded[i] = (float64(p) / float64(params.Q)) * float64(y[i])
	}

	res := make([]int16, len(y))
	for i := 0; i < len(unrounded); i++ {
		res[i] = int16(math.Round(unrounded[i]))
	}

	var out kyberk2so.Poly
	copy(out[:], res)

	out = kyberk2so.PolyReduce(out)
	return y
}
