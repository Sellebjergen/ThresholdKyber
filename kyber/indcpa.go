// indcpa.go - Kyber IND-CPA encryption.
//
// To the extent possible under law, Yawning Angel has waived all copyright
// and related or neighboring rights to the software, using the Creative
// Commons "CC0" public domain dedication. See LICENSE or
// <http://creativecommons.org/publicdomain/zero/1.0/> for full details.

package kyber

import (
	"io"

	"golang.org/x/crypto/sha3"
)

// Serialize the public key as concatenation of the compressed and serialized
// vector of polynomials pk and the public seed used to generate the matrix A.
func packPublicKey(r []byte, pk *PolyVec, seed []byte) {
	pk.compress(r)
	copy(r[pk.compressedSize():], seed[:SymSize])
}

// De-serialize and decompress public key from a byte array; approximate
// inverse of packPublicKey.
func unpackPublicKey(pk *PolyVec, seed, packedPk []byte) {
	pk.decompress(packedPk)

	off := pk.compressedSize()
	copy(seed, packedPk[off:off+SymSize])
}

// Serialize the ciphertext as concatenation of the compressed and serialized
// vector of polynomials b and the compressed and serialized polynomial v.
func packCiphertext(r []byte, b *PolyVec, v *Poly) {
	b.compress(r)
	v.Compress(r[b.compressedSize():])
}

// De-serialize and decompress ciphertext from a byte array; approximate
// inverse of packCiphertext.
func UnpackCiphertext(b *PolyVec, v *Poly, c []byte) {
	b.decompress(c)
	v.decompress(c[b.compressedSize():])
}

// Serialize the secret key.
func PackSecretKey(r []byte, sk *PolyVec) {
	sk.toBytes(r)
}

// De-serialize the secret key; inverse of PackSecretKey.
func UnpackSecretKey(sk *PolyVec, packedSk []byte) {
	sk.fromBytes(packedSk)
}

// Deterministically generate matrix A (or the transpose of A) from a seed.
// Entries of the matrix are polynomials that look uniformly random. Performs
// rejection sampling on output of SHAKE-128.
func genMatrix(a []PolyVec, seed []byte, transposed bool) {
	const (
		shake128Rate = 168 // xof.BlockSize() is not a constant.
		maxBlocks    = 4
	)
	var buf [shake128Rate * maxBlocks]byte

	var extSeed [SymSize + 2]byte
	copy(extSeed[:SymSize], seed)

	xof := sha3.NewShake128()

	for i, v := range a {
		for j, p := range v.Vec {
			if transposed {
				extSeed[SymSize] = byte(i)
				extSeed[SymSize+1] = byte(j)
			} else {
				extSeed[SymSize] = byte(j)
				extSeed[SymSize+1] = byte(i)
			}

			xof.Write(extSeed[:])
			xof.Read(buf[:])

			for ctr, pos, maxPos := 0, 0, len(buf); ctr < kyberN; {
				val := (uint16(buf[pos]) | (uint16(buf[pos+1]) << 8)) & 0x1fff
				if val < kyberQ {
					p.Coeffs[ctr] = val
					ctr++
				}
				if pos += 2; pos == maxPos {
					// On the unlikely chance 4 blocks is insufficient,
					// incrementally squeeze out 1 block at a time.
					xof.Read(buf[:shake128Rate])
					pos, maxPos = 0, shake128Rate
				}
			}

			xof.Reset()
		}
	}
}

type IndcpaPublicKey struct {
	packed []byte
	h      [32]byte
}

func (pk *IndcpaPublicKey) toBytes() []byte {
	return pk.packed
}

func (pk *IndcpaPublicKey) fromBytes(p *ParameterSet, b []byte) error {
	if len(b) != p.indcpaPublicKeySize {
		return ErrInvalidKeySize
	}

	pk.packed = make([]byte, len(b))
	copy(pk.packed, b)
	pk.h = sha3.Sum256(b)

	return nil
}

type IndcpaSecretKey struct {
	Packed []byte
}

func (sk *IndcpaSecretKey) fromBytes(p *ParameterSet, b []byte) error {
	if len(b) != p.IndcpaSecretKeySize {
		return ErrInvalidKeySize
	}

	sk.Packed = make([]byte, len(b))
	copy(sk.Packed, b)

	return nil
}

// Generates public and private key for the CPA-secure public-key encryption
// scheme underlying Kyber.
func (p *ParameterSet) IndcpaKeyPair(rng io.Reader) (*IndcpaPublicKey, *IndcpaSecretKey, error) {
	buf := make([]byte, SymSize+SymSize)
	if _, err := io.ReadFull(rng, buf[:SymSize]); err != nil {
		return nil, nil, err
	}

	sk := &IndcpaSecretKey{
		Packed: make([]byte, p.IndcpaSecretKeySize),
	}
	pk := &IndcpaPublicKey{
		packed: make([]byte, p.indcpaPublicKeySize),
	}

	h := sha3.New512()
	h.Write(buf[:SymSize])
	buf = buf[:0] // Reuse the backing store.
	buf = h.Sum(buf)
	publicSeed, noiseSeed := buf[:SymSize], buf[SymSize:]

	a := p.allocMatrix()
	genMatrix(a, publicSeed, false)

	var nonce byte
	skpv := p.AllocPolyVec()
	for _, pv := range skpv.Vec {
		pv.getNoise(noiseSeed, nonce, p.eta)
		nonce++
	}

	skpv.ntt()

	e := p.AllocPolyVec()
	for _, pv := range e.Vec {
		pv.getNoise(noiseSeed, nonce, p.eta)
		nonce++
	}

	// matrix-vector multiplication
	pkpv := p.AllocPolyVec()
	for i, pv := range pkpv.Vec {
		pv.pointwiseAcc(&skpv, &a[i])
	}

	pkpv.invntt()
	pkpv.Add(&pkpv, &e)

	PackSecretKey(sk.Packed, &skpv)
	packPublicKey(pk.packed, &pkpv, publicSeed)
	pk.h = sha3.Sum256(pk.packed)

	return pk, sk, nil
}

// Encryption function of the CPA-secure public-key encryption scheme
// underlying Kyber.
func (p *ParameterSet) IndcpaEncrypt(c, m []byte, pk *IndcpaPublicKey, coins []byte) {
	var k, v, epp Poly
	var seed [SymSize]byte

	pkpv := p.AllocPolyVec()
	unpackPublicKey(&pkpv, seed[:], pk.packed)

	k.FromMsg(m)

	pkpv.ntt()

	// A
	at := p.allocMatrix()
	genMatrix(at, seed[:], true)

	// s
	var nonce byte
	sp := p.AllocPolyVec()
	for _, pv := range sp.Vec {
		pv.getNoise(coins, nonce, p.eta)
		nonce++
	}

	sp.ntt()

	// e
	ep := p.AllocPolyVec()
	for _, pv := range ep.Vec {
		pv.getNoise(coins, nonce, p.eta)
		nonce++
	}

	// matrix-vector multiplication
	// A * s
	bp := p.AllocPolyVec()
	for i, pv := range bp.Vec {
		pv.pointwiseAcc(&sp, &at[i])
	}

	bp.invntt()
	// (A*s) + e
	bp.Add(&bp, &ep)

	v.pointwiseAcc(&pkpv, &sp)
	v.Invntt()

	epp.getNoise(coins, nonce, p.eta) // Don't need to increment nonce.

	v.add(&v, &epp)
	v.add(&v, &k)

	packCiphertext(c, &bp, &v)
}

// Decryption function of the CPA-secure public-key encryption scheme
// underlying Kyber.
func (p *ParameterSet) IndcpaDecrypt(m, c []byte, sk *IndcpaSecretKey) {
	var v, mp Poly

	skpv, bp := p.AllocPolyVec(), p.AllocPolyVec()
	UnpackCiphertext(&bp, &v, c)
	UnpackSecretKey(&skpv, sk.Packed)

	bp.ntt()

	mp.pointwiseAcc(&skpv, &bp)
	mp.Invntt()

	mp.sub(&mp, &v)

	mp.ToMsg(m)
}

func (p *ParameterSet) allocMatrix() []PolyVec {
	m := make([]PolyVec, 0, p.k)
	for i := 0; i < p.k; i++ {
		m = append(m, p.AllocPolyVec())
	}
	return m
}

func (p *ParameterSet) AllocPolyVec() PolyVec {
	vec := make([]*Poly, 0, p.k)
	for i := 0; i < p.k; i++ {
		vec = append(vec, new(Poly))
	}

	return PolyVec{vec}
}
