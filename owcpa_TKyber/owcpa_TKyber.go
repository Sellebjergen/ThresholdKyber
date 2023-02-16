package owcpa_TKyber

import (
	"crypto/rand"
	"math"

	kyberk2so "ThresholdKyber.com/m/kyber-k2so"
)

func Setup(params *OwcpaParams, n int, t int) ([]byte, [][]kyberk2so.PolyVec) {
	// Run setup to get Kyber KeyPair
	sk, pk, _ := kyberk2so.IndcpaKeypair(kyberk2so.ParamsK)
	sk_unpacked := kyberk2so.IndcpaUnpackPrivateKey(sk, kyberk2so.ParamsK)

	// Perform secret sharing
	sk_shares := params.LSS_scheme.Share(sk_unpacked, n, t)

	return pk, sk_shares
}

func Enc(params *OwcpaParams, msg []byte, pk []byte) []byte {
	coins := make([]byte, 32)
	rand.Read(coins)
	ct, _ := kyberk2so.IndcpaEncrypt(msg, pk, coins, kyberk2so.ParamsK)
	return ct
}

func PartDec(params *OwcpaParams, sk_i []kyberk2so.PolyVec, ct []byte, party int) []kyberk2so.Poly {
	var zero kyberk2so.Poly
	// Sample noise
	e_i := params.D_flood_dist.SampleNoise(params.Q, 255, params.Sigma) // TODO: Fix params

	u, v := kyberk2so.IndcpaUnpackCiphertext(ct, kyberk2so.ParamsK)

	// Inner prod
	L := len(sk_i)
	d_i := make([]kyberk2so.Poly, L)

	kyberk2so.PolyvecNtt(u, kyberk2so.ParamsK)

	for j := 0; j < L; j++ {
		d_i[j] = kyberk2so.PolyvecPointWiseAccMontgomery(sk_i[j], u, kyberk2so.ParamsK)
		d_i[j] = kyberk2so.PolyInvNttToMont(d_i[j])
		if party == 0 {
			d_i[j] = kyberk2so.PolySub(v, d_i[j])
		} else {
			d_i[j] = kyberk2so.PolySub(zero, d_i[j])
		}

		// Add noise
		d_i[j] = kyberk2so.PolyAdd(d_i[j], e_i)

		d_i[j] = kyberk2so.PolyReduce(d_i[j])
	}

	return d_i
}

func Combine(params *OwcpaParams, ct []byte, d_is [][]kyberk2so.Poly, n int, t int) kyberk2so.Poly {
	y := params.LSS_scheme.Rec(d_is, n, t)
	return y
}

func Downscale(in kyberk2so.Poly, p int, q int) kyberk2so.Poly {
	y := kyberk2so.PolyReduce(in)
	unrounded := make([]float64, len(in))

	for i := 0; i < len(unrounded); i++ {
		unrounded[i] = (float64(p) / float64(q)) * float64(y[i])
	}

	res := make([]int16, len(in))
	for i := 0; i < len(unrounded); i++ {
		res[i] = int16(math.Round(unrounded[i]))
	}

	mod2 := make([]int16, len(in))
	for i := 0; i < len(mod2); i++ {
		mod2[i] = res[i] % 2
	}

	var out kyberk2so.Poly
	copy(out[:], mod2)
	return out
}

func Upscale(in kyberk2so.Poly, p int, q int) kyberk2so.Poly {
	unrounded := make([]float64, len(in))
	for i := 0; i < len(unrounded); i++ {
		unrounded[i] = (float64(q) / float64(p))
	}

	factor := make([]int16, len(in))
	for i := 0; i < len(factor); i++ {
		factor[i] = int16(math.Round(unrounded[i]))
	}

	scaled := make([]int16, len(in))
	for i := 0; i < len(scaled); i++ {
		scaled[i] = factor[i] * in[i]
	}

	var out kyberk2so.Poly
	copy(out[:], scaled)
	return out
}
