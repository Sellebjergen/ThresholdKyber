package owcpa_TKyber

import (
	"crypto/rand"

	"ThresholdKyber.com/m/kyber"
)

func Setup(params *OwcpaParams, n int, t int) (*kyber.IndcpaPublicKey, []kyber.PolyVec) {
	// Run setup to get Kyber KeyPair
	pk, sk, _ := params.KyberParams.IndcpaKeyPair(rand.Reader)
	polyVec_sk := params.KyberParams.AllocPolyVec()
	kyber.UnpackSecretKey(&polyVec_sk, sk.Packed)

	// Perform secret sharing
	sk_shares := params.LSS_scheme.Share(polyVec_sk, n)

	return pk, sk_shares
}

func Enc(params *OwcpaParams, msg []byte, pk *kyber.IndcpaPublicKey) []byte {
	coins := make([]byte, 32)
	rand.Read(coins)

	ct := make([]byte, params.KyberParams.CipherTextSize())
	kyber.Kyber512.IndcpaEncrypt(ct, msg, pk, coins)

	return ct
}

func PartDec(params *OwcpaParams, sk_i kyber.PolyVec, ct []byte, party int) *kyber.Poly {
	var v, d_i, zero kyber.Poly
	// Sample noise
	e_i := params.D_flood_dist.SampleNoise(params.Q, 255, params.Sigma) // TODO: Fix params

	bp := params.KyberParams.AllocPolyVec()
	kyber.UnpackCiphertext(&bp, &v, ct) // This will be NTT form.

	// Inner prod
	bp.Ntt()
	d_i.PointwiseAcc(&sk_i, &bp)
	d_i.Invntt()

	if party == 0 {
		d_i.Sub(&v, &d_i)
	} else {
		d_i.Sub(&zero, &d_i)
	}

	// Add noise
	d_i.Add(&d_i, e_i)

	return &d_i
}

func Combine(params *OwcpaParams, ct []byte, d_is []*kyber.Poly) *kyber.Poly {
	/* p := 2 */

	y := params.LSS_scheme.Rec(d_is)
	/* unrounded := make([]float64, len(y.Coeffs))

	for i := 0; i < len(unrounded); i++ {
		unrounded[i] = (float64(p) / float64(rq.q)) * float64(y.Coeffs[i])
	}

	res := make([]uint16, len(y.Coeffs))
	for i := 0; i < len(unrounded); i++ {
		res[i] = uint16(math.Round(unrounded[i]))
	}
	out := new(kyber.Poly)
	copy(out.Coeffs[:], res)

	for i := 0; i < len(out.Coeffs); i++ {
		out.Coeffs[i] = kyber.Freeze(out.Coeffs[i])
	} */

	return y
}
