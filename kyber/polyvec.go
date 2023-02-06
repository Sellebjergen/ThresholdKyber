// polyvec.go - Vector of Kyber polynomials.
//
// To the extent possible under law, Yawning Angel has waived all copyright
// and related or neighboring rights to the software, using the Creative
// Commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package kyber

type PolyVec struct {
	Vec []*Poly
}

// Compress and serialize vector of polynomials.
func (v *PolyVec) compress(r []byte) {
	var off int
	for _, vec := range v.Vec {
		for j := 0; j < kyberN/8; j++ {
			var t [8]uint16
			for k := 0; k < 8; k++ {
				t[k] = uint16((((uint32(Freeze(vec.Coeffs[8*j+k])) << 11) + kyberQ/2) / kyberQ) & 0x7ff)
			}

			r[off+11*j+0] = byte(t[0] & 0xff)
			r[off+11*j+1] = byte((t[0] >> 8) | ((t[1] & 0x1f) << 3))
			r[off+11*j+2] = byte((t[1] >> 5) | ((t[2] & 0x03) << 6))
			r[off+11*j+3] = byte((t[2] >> 2) & 0xff)
			r[off+11*j+4] = byte((t[2] >> 10) | ((t[3] & 0x7f) << 1))
			r[off+11*j+5] = byte((t[3] >> 7) | ((t[4] & 0x0f) << 4))
			r[off+11*j+6] = byte((t[4] >> 4) | ((t[5] & 0x01) << 7))
			r[off+11*j+7] = byte((t[5] >> 1) & 0xff)
			r[off+11*j+8] = byte((t[5] >> 9) | ((t[6] & 0x3f) << 2))
			r[off+11*j+9] = byte((t[6] >> 6) | ((t[7] & 0x07) << 5))
			r[off+11*j+10] = byte((t[7] >> 3))
		}
		off += compressedCoeffSize
	}
}

// De-serialize and decompress vector of polynomials; approximate inverse of
// polyVec.compress().
func (v *PolyVec) decompress(a []byte) {
	var off int
	for _, vec := range v.Vec {
		for j := 0; j < kyberN/8; j++ {
			vec.Coeffs[8*j+0] = uint16((((uint32(a[off+11*j+0]) | ((uint32(a[off+11*j+1]) & 0x07) << 8)) * kyberQ) + 1024) >> 11)
			vec.Coeffs[8*j+1] = uint16(((((uint32(a[off+11*j+1]) >> 3) | ((uint32(a[off+11*j+2]) & 0x3f) << 5)) * kyberQ) + 1024) >> 11)
			vec.Coeffs[8*j+2] = uint16(((((uint32(a[off+11*j+2]) >> 6) | ((uint32(a[off+11*j+3]) & 0xff) << 2) | ((uint32(a[off+11*j+4]) & 0x01) << 10)) * kyberQ) + 1024) >> 11)
			vec.Coeffs[8*j+3] = uint16(((((uint32(a[off+11*j+4]) >> 1) | ((uint32(a[off+11*j+5]) & 0x0f) << 7)) * kyberQ) + 1024) >> 11)
			vec.Coeffs[8*j+4] = uint16(((((uint32(a[off+11*j+5]) >> 4) | ((uint32(a[off+11*j+6]) & 0x7f) << 4)) * kyberQ) + 1024) >> 11)
			vec.Coeffs[8*j+5] = uint16(((((uint32(a[off+11*j+6]) >> 7) | ((uint32(a[off+11*j+7]) & 0xff) << 1) | ((uint32(a[off+11*j+8]) & 0x03) << 9)) * kyberQ) + 1024) >> 11)
			vec.Coeffs[8*j+6] = uint16(((((uint32(a[off+11*j+8]) >> 2) | ((uint32(a[off+11*j+9]) & 0x1f) << 6)) * kyberQ) + 1024) >> 11)
			vec.Coeffs[8*j+7] = uint16(((((uint32(a[off+11*j+9]) >> 5) | ((uint32(a[off+11*j+10]) & 0xff) << 3)) * kyberQ) + 1024) >> 11)
		}
		off += compressedCoeffSize
	}
}

// Serialize vector of polynomials.
func (v *PolyVec) toBytes(r []byte) {
	for i, p := range v.Vec {
		p.ToBytes(r[i*polySize:])
	}
}

// De-serialize vector of polynomials; inverse of polyVec.toBytes().
func (v *PolyVec) fromBytes(a []byte) {
	for i, p := range v.Vec {
		p.fromBytes(a[i*polySize:])
	}
}

// Apply forward NTT to all elements of a vector of polynomials.
func (v *PolyVec) Ntt() {
	for _, p := range v.Vec {
		p.Ntt()
	}
}

// Apply inverse NTT to all elements of a vector of polynomials.
func (v *PolyVec) Invntt() {
	for _, p := range v.Vec {
		p.Invntt()
	}
}

// Pointwise multiply elements of a and b and accumulate into p.
func (p *Poly) PointwiseAcc(a, b *PolyVec) {
	hardwareAccelImpl.pointwiseAccFn(p, a, b)
}

// Add vectors of polynomials.
func (v *PolyVec) Add(a, b *PolyVec) {
	for i, p := range v.Vec {
		p.Add(a.Vec[i], b.Vec[i])
	}
}

// Sub vectors of polynomials.
func (v *PolyVec) Sub(a, b *PolyVec) {
	for i, p := range v.Vec {
		p.Sub(a.Vec[i], b.Vec[i])
	}
}

// Get compressed and serialized size in bytes.
func (v *PolyVec) compressedSize() int {
	return len(v.Vec) * compressedCoeffSize
}

func pointwiseAccRef(p *Poly, a, b *PolyVec) {
	for j := 0; j < kyberN; j++ {
		t := montgomeryReduce(4613 * uint32(b.Vec[0].Coeffs[j])) // 4613 = 2^{2*18} % q
		p.Coeffs[j] = montgomeryReduce(uint32(a.Vec[0].Coeffs[j]) * uint32(t))
		for i := 1; i < len(a.Vec); i++ { // len(a.vec) == kyberK
			t = montgomeryReduce(4613 * uint32(b.Vec[i].Coeffs[j]))
			p.Coeffs[j] += montgomeryReduce(uint32(a.Vec[i].Coeffs[j]) * uint32(t))
		}

		p.Coeffs[j] = barrettReduce(p.Coeffs[j])
	}
}
