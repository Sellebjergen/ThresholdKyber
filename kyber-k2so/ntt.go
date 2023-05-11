/* SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: MIT */

package kyberk2so

var nttZetas [128]int16 = [128]int16{
	1018, 10530, 7065, 4188, 357, 376, 8340, 7067, 10328, 6209, 6969, 3364, 10023, 2695, 5898, 2236, 1931, 1341, 10412, 6930, 8124, 5213, 4095, 5578, 5690, 4035, 5878, 9233, 918, 2503, 3012, 4347, 6651, 6011, 6009, 5877, 9906, 4379, 2984, 7691, 8177, 6564, 544, 1085, 9744, 8365, 3085, 9454, 268, 9469, 3550, 4, 4129, 7662, 2935, 2790, 205, 6240, 7771, 1287, 8575, 8188, 6137, 10560, 8016, 4286, 578, 3169, 10353, 151, 5294, 636, 9181, 8013, 8547, 779, 336, 2884, 9747, 326, 2234, 6630, 864, 7416, 9783, 635, 472, 467, 3715, 9485, 9167, 9685, 5116, 8069, 1324, 7780, 10078, 2271, 2213, 9138, 7696, 8708, 512, 7979, 567, 7555, 5023, 9959, 2998, 6019, 3441, 5341, 7772, 5776, 6173, 116, 1458, 7138, 9844, 5639, 5945, 8912, 10299, 5064, 1356, 886, 5925, 5156,
}

const f = 5072

// nttFqMul performs multiplication followed by Montgomery reduction
// and returns a 16-bit integer congruent to `a*b*R^{-1} mod Q`.
func nttFqMul(a int16, b int16) int16 {
	return ByteopsMontgomeryReduce(int32(a) * int32(b))
}

// ntt performs an inplace number-theoretic transform (NTT) in `Rq`.
// The input is in standard order, the output is in bit-reversed order.
func ntt(r Poly) Poly {
	j := 0
	k := 1
	for l := 128; l >= 2; l >>= 1 {
		for start := 0; start < 256; start = j + l {
			zeta := nttZetas[k]
			k = k + 1
			for j = start; j < start+l; j++ {
				t := nttFqMul(zeta, r[j+l])
				r[j+l] = r[j] - t
				r[j] = r[j] + t
			}
		}
	}
	return r
}

// nttInv performs an inplace inverse number-theoretic transform (NTT)
// in `Rq` and multiplication by Montgomery factor 2^16.
// The input is in bit-reversed order, the output is in standard order.
func nttInv(r Poly) Poly {
	j := 0
	k := 127
	for l := 2; l <= 128; l <<= 1 {
		for start := 0; start < 256; start = j + l {
			zeta := nttZetas[k]
			k = k - 1
			for j = start; j < start+l; j++ {
				t := r[j]
				r[j] = ByteopsBarrettReduce(t + r[j+l])
				r[j+l] = r[j+l] - t
				r[j+l] = nttFqMul(zeta, r[j+l])
			}
		}
	}
	for j := 0; j < 256; j++ {
		r[j] = nttFqMul(r[j], f)
	}
	return r
}

// nttBaseMul performs the multiplication of polynomials
// in `Zq[X]/(X^2-zeta)`. Used for multiplication of elements
// in `Rq` in the number-theoretic transformation domain.
func nttBaseMul(
	a0 int16, a1 int16,
	b0 int16, b1 int16,
	zeta int16,
) [2]int16 {
	var r [2]int16
	r[0] = nttFqMul(a1, b1)
	r[0] = nttFqMul(r[0], zeta)
	r[0] = r[0] + nttFqMul(a0, b0)
	r[1] = nttFqMul(a0, b1)
	r[1] = r[1] + nttFqMul(a1, b0)
	return r
}
