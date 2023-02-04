// Poly.go - Kyber polynomial.
//
// To the extent possible under law, Yawning Angel has waived all copyright
// and related or neighboring rights to the software, using the Creative
// Commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package kyber

import "golang.org/x/crypto/sha3"

// Elements of R_q = Z_q[X]/(X^n + 1). Represents polynomial Coeffs[0] +
// X*Coeffs[1] + X^2*xoeffs[2] + ... + X^{n-1}*Coeffs[n-1].
type Poly struct {
	Coeffs [kyberN]uint16
}

// Compression and subsequent serialization of a polynomial.
func (p *Poly) Compress(r []byte) {
	var t [8]uint32

	for i, k := 0, 0; i < kyberN; i, k = i+8, k+3 {
		for j := 0; j < 8; j++ {
			t[j] = uint32((((freeze(p.Coeffs[i+j]) << 3) + kyberQ/2) / kyberQ) & 7)
		}

		r[k] = byte(t[0] | (t[1] << 3) | (t[2] << 6))
		r[k+1] = byte((t[2] >> 2) | (t[3] << 1) | (t[4] << 4) | (t[5] << 7))
		r[k+2] = byte((t[5] >> 1) | (t[6] << 2) | (t[7] << 5))
	}
}

// De-serialization and subsequent decompression of a polynomial; approximate
// inverse of Poly.compress().
func (p *Poly) decompress(a []byte) {
	for i, off := 0, 0; i < kyberN; i, off = i+8, off+3 {
		p.Coeffs[i+0] = ((uint16(a[off]&7) * kyberQ) + 4) >> 3
		p.Coeffs[i+1] = (((uint16(a[off]>>3) & 7) * kyberQ) + 4) >> 3
		p.Coeffs[i+2] = (((uint16(a[off]>>6) | (uint16(a[off+1]<<2) & 4)) * kyberQ) + 4) >> 3
		p.Coeffs[i+3] = (((uint16(a[off+1]>>1) & 7) * kyberQ) + 4) >> 3
		p.Coeffs[i+4] = (((uint16(a[off+1]>>4) & 7) * kyberQ) + 4) >> 3
		p.Coeffs[i+5] = (((uint16(a[off+1]>>7) | (uint16(a[off+2]<<1) & 6)) * kyberQ) + 4) >> 3
		p.Coeffs[i+6] = (((uint16(a[off+2]>>2) & 7) * kyberQ) + 4) >> 3
		p.Coeffs[i+7] = (((uint16(a[off+2] >> 5)) * kyberQ) + 4) >> 3
	}
}

// Serialization of a polynomial.
func (p *Poly) ToBytes(r []byte) {
	var t [8]uint16

	for i := 0; i < kyberN/8; i++ {
		for j := 0; j < 8; j++ {
			t[j] = freeze(p.Coeffs[8*i+j])
		}

		r[13*i+0] = byte(t[0] & 0xff)
		r[13*i+1] = byte((t[0] >> 8) | ((t[1] & 0x07) << 5))
		r[13*i+2] = byte((t[1] >> 3) & 0xff)
		r[13*i+3] = byte((t[1] >> 11) | ((t[2] & 0x3f) << 2))
		r[13*i+4] = byte((t[2] >> 6) | ((t[3] & 0x01) << 7))
		r[13*i+5] = byte((t[3] >> 1) & 0xff)
		r[13*i+6] = byte((t[3] >> 9) | ((t[4] & 0x0f) << 4))
		r[13*i+7] = byte((t[4] >> 4) & 0xff)
		r[13*i+8] = byte((t[4] >> 12) | ((t[5] & 0x7f) << 1))
		r[13*i+9] = byte((t[5] >> 7) | ((t[6] & 0x03) << 6))
		r[13*i+10] = byte((t[6] >> 2) & 0xff)
		r[13*i+11] = byte((t[6] >> 10) | ((t[7] & 0x1f) << 3))
		r[13*i+12] = byte(t[7] >> 5)
	}
}

// De-serialization of a polynomial; inverse of Poly.toBytes().
func (p *Poly) fromBytes(a []byte) {
	for i := 0; i < kyberN/8; i++ {
		p.Coeffs[8*i+0] = uint16(a[13*i+0]) | ((uint16(a[13*i+1]) & 0x1f) << 8)
		p.Coeffs[8*i+1] = (uint16(a[13*i+1]) >> 5) | (uint16(a[13*i+2]) << 3) | ((uint16(a[13*i+3]) & 0x03) << 11)
		p.Coeffs[8*i+2] = (uint16(a[13*i+3]) >> 2) | ((uint16(a[13*i+4]) & 0x7f) << 6)
		p.Coeffs[8*i+3] = (uint16(a[13*i+4]) >> 7) | (uint16(a[13*i+5]) << 1) | ((uint16(a[13*i+6]) & 0x0f) << 9)
		p.Coeffs[8*i+4] = (uint16(a[13*i+6]) >> 4) | (uint16(a[13*i+7]) << 4) | ((uint16(a[13*i+8]) & 0x01) << 12)
		p.Coeffs[8*i+5] = (uint16(a[13*i+8]) >> 1) | ((uint16(a[13*i+9]) & 0x3f) << 7)
		p.Coeffs[8*i+6] = (uint16(a[13*i+9]) >> 6) | (uint16(a[13*i+10]) << 2) | ((uint16(a[13*i+11]) & 0x07) << 10)
		p.Coeffs[8*i+7] = (uint16(a[13*i+11]) >> 3) | (uint16(a[13*i+12]) << 5)
	}
}

// Convert 32-byte message to polynomial.
func (p *Poly) FromMsg(msg []byte) {
	for i, v := range msg[:SymSize] {
		for j := 0; j < 8; j++ {
			mask := -((uint16(v) >> uint(j)) & 1)
			p.Coeffs[8*i+j] = mask & ((kyberQ + 1) / 2)
		}
	}
}

// Convert polynomial to 32-byte message.
func (p *Poly) ToMsg(msg []byte) {
	for i := 0; i < SymSize; i++ {
		msg[i] = 0
		for j := 0; j < 8; j++ {
			t := (((freeze(p.Coeffs[8*i+j]) << 1) + kyberQ/2) / kyberQ) & 1
			msg[i] |= byte(t << uint(j))
		}
	}
}

// Sample a polynomial deterministically from a seed and a nonce, with output
// polynomial close to centered binomial distribution with parameter eta.
func (p *Poly) getNoise(seed []byte, nonce byte, eta int) {
	extSeed := make([]byte, 0, SymSize+1)
	extSeed = append(extSeed, seed...)
	extSeed = append(extSeed, nonce)

	buf := make([]byte, eta*kyberN/4)
	sha3.ShakeSum256(buf, extSeed)

	p.cbd(buf, eta)
}

// Computes negacyclic number-theoretic transform (NTT) of a polynomial in
// place; inputs assumed to be in normal order, output in bitreversed order.
func (p *Poly) ntt() {
	hardwareAccelImpl.nttFn(&p.Coeffs)
}

// Computes inverse of negacyclic number-theoretic transform (NTT) of a
// polynomial in place; inputs assumed to be in bitreversed order, output in
// normal order.
func (p *Poly) invntt() {
	hardwareAccelImpl.invnttFn(&p.Coeffs)
}

// Add two polynomials.
func (p *Poly) add(a, b *Poly) {
	for i := range p.Coeffs {
		p.Coeffs[i] = barrettReduce(a.Coeffs[i] + b.Coeffs[i])
	}
}

// Subtract two polynomials.
func (p *Poly) sub(a, b *Poly) {
	for i := range p.Coeffs {
		p.Coeffs[i] = barrettReduce(a.Coeffs[i] + 3*kyberQ - b.Coeffs[i])
	}
}

// todo: maybe move?
func (p *Poly) GetDegree() int {
	i := len(p.Coeffs)
	for i > 0 && p.Coeffs[i] == 0 {
		i--
	}
	return i
}
