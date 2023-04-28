/* SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: MIT */

package kyberk2so

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"

	"golang.org/x/crypto/sha3"
)

// IndcpaPackPublicKey serializes the public key as a concatenation of the
// serialized vector of polynomials of the public key, and the public seed
// used to generate the matrix `A`.
func IndcpaPackPublicKey(publicKey PolyVec, seed []byte, paramsK int) []byte {
	return append(polyvecToBytes(publicKey, paramsK), seed...)
}

// IndcpaUnpackPublicKey de-serializes the public key from a byte array
// and represents the approximate inverse of indcpaPackPublicKey.
func IndcpaUnpackPublicKey(packedPublicKey []byte, paramsK int) (PolyVec, []byte) {
	switch paramsK {
	case 2:
		publicKeyPolyvec := polyvecFromBytes(packedPublicKey[:paramsPolyvecBytesK512], paramsK)
		seed := packedPublicKey[paramsPolyvecBytesK512:]
		return publicKeyPolyvec, seed
	case 3:
		publicKeyPolyvec := polyvecFromBytes(packedPublicKey[:paramsPolyvecBytesK768], paramsK)
		seed := packedPublicKey[paramsPolyvecBytesK768:]
		return publicKeyPolyvec, seed
	case 4:
		publicKeyPolyvec := polyvecFromBytes(packedPublicKey[:paramsPolyvecBytesK1024], paramsK)
		seed := packedPublicKey[paramsPolyvecBytesK1024:]
		return publicKeyPolyvec, seed
	case 5:
		publicKeyPolyvec := polyvecFromBytes(packedPublicKey[:paramsPolyvecBytesK1280], paramsK)
		seed := packedPublicKey[paramsPolyvecBytesK1280:]
		return publicKeyPolyvec, seed
	case 6:
		publicKeyPolyvec := polyvecFromBytes(packedPublicKey[:paramsPolyvecBytesK1536], paramsK)
		seed := packedPublicKey[paramsPolyvecBytesK1536:]
		return publicKeyPolyvec, seed
	default:
		publicKeyPolyvec := polyvecFromBytes(packedPublicKey[:paramsPolyvecBytesK1792], paramsK)
		seed := packedPublicKey[paramsPolyvecBytesK1792:]
		return publicKeyPolyvec, seed
	}
}

// IndcpaPackPrivateKey serializes the private key.
func IndcpaPackPrivateKey(privateKey PolyVec, paramsK int) []byte {
	return polyvecToBytes(privateKey, paramsK)
}

// IndcpaUnpackPrivateKey de-serializes the private key and represents
// the inverse of indcpaPackPrivateKey.
func IndcpaUnpackPrivateKey(packedPrivateKey []byte, paramsK int) PolyVec {
	return polyvecFromBytes(packedPrivateKey, paramsK)
}

// IndcpaPackCiphertext serializes the ciphertext as a concatenation of
// the compressed and serialized vector of polynomials `b` and the
// compressed and serialized polynomial `v`.
func IndcpaPackCiphertext(b PolyVec, v Poly, paramsK int) []byte {
	return append(polyvecCompress(b, paramsK), PolyCompress(v, paramsK)...)
}

// IndcpaUnpackCiphertext de-serializes and decompresses the ciphertext
// from a byte array, and represents the approximate inverse of
// indcpaPackCiphertext.
func IndcpaUnpackCiphertext(c []byte, paramsK int) (PolyVec, Poly) {
	switch paramsK {
	case 2:
		b := polyvecDecompress(c[:paramsPolyvecCompressedBytesK512], paramsK)
		v := polyDecompress(c[paramsPolyvecCompressedBytesK512:], paramsK)
		return b, v
	case 3:
		b := polyvecDecompress(c[:paramsPolyvecCompressedBytesK768], paramsK)
		v := polyDecompress(c[paramsPolyvecCompressedBytesK768:], paramsK)
		return b, v
	case 4:
		b := polyvecDecompress(c[:paramsPolyvecCompressedBytesK1024], paramsK)
		v := polyDecompress(c[paramsPolyvecCompressedBytesK1024:], paramsK)
		return b, v
	case 5:
		b := polyvecDecompress(c[:paramsPolyvecCompressedBytesK1280], paramsK)
		v := polyDecompress(c[paramsPolyvecCompressedBytesK1280:], paramsK)
		return b, v
	case 6:
		b := polyvecDecompress(c[:paramsPolyvecCompressedBytesK1536], paramsK)
		v := polyDecompress(c[paramsPolyvecCompressedBytesK1536:], paramsK)
		return b, v
	default:
		b := polyvecDecompress(c[:paramsPolyvecCompressedBytesK1792], paramsK)
		v := polyDecompress(c[paramsPolyvecCompressedBytesK1792:], paramsK)
		return b, v
	}
}

// indcpaRejUniform runs rejection sampling on uniform random bytes
// to generate uniform random integers modulo `Q`.
func indcpaRejUniform(buf []byte, bufl int, l int) (Poly, int) {
	var r Poly
	var d1 uint16
	var d2 uint16
	i := 0
	j := 0
	for i < l && j+3 <= bufl {
		d1 = (uint16((buf[j])>>0) | (uint16(buf[j+1]) << 8)) & 0xFFF
		d2 = (uint16((buf[j+1])>>4) | (uint16(buf[j+2]) << 4)) & 0xFFF
		j = j + 3
		if d1 < uint16(ParamsQ) {
			r[i] = int16(d1)
			i = i + 1
		}
		if i < l && d2 < uint16(ParamsQ) {
			r[i] = int16(d2)
			i = i + 1
		}
	}
	return r, i
}

// IndcpaGenMatrix deterministically generates a matrix `A` (or the transpose of `A`)
// from a seed. Entries of the matrix are polynomials that look uniformly random.
// Performs rejection sampling on the output of an extendable-output function (XOF).
func IndcpaGenMatrix(seed []byte, transposed bool, paramsK int) ([]PolyVec, error) {
	r := make([]PolyVec, paramsK)
	buf := make([]byte, 672)
	xof := sha3.NewShake128()
	ctr := 0
	for i := 0; i < paramsK; i++ {
		r[i] = PolyvecNew(paramsK)
		for j := 0; j < paramsK; j++ {
			xof.Reset()
			var err error
			if transposed {
				_, err = xof.Write(append(seed, []byte{byte(i), byte(j)}...))
			} else {
				_, err = xof.Write(append(seed, []byte{byte(j), byte(i)}...))
			}
			if err != nil {
				return []PolyVec{}, err
			}
			_, err = xof.Read(buf)
			if err != nil {
				return []PolyVec{}, err
			}
			r[i][j], ctr = indcpaRejUniform(buf[:504], 504, ParamsN)
			for ctr < ParamsN {
				missing, ctrn := indcpaRejUniform(buf[504:672], 168, ParamsN-ctr)
				for k := ctr; k < ParamsN; k++ {
					r[i][j][k] = missing[k-ctr]
				}
				ctr = ctr + ctrn
			}
		}
	}
	return r, nil
}

// indcpaPrf provides a pseudo-random function (PRF) which returns
// a byte array of length `l`, using the provided key and nonce
// to instantiate the PRF's underlying hash function.
func indcpaPrf(l int, key []byte, nonce byte) []byte {
	hash := make([]byte, l)
	fmt.Println(l)
	sha3.ShakeSum256(hash, append(key, nonce))
	fmt.Println(hex.EncodeToString(hash))
	return hash
}

// IndcpaKeypair generates public and private keys for the CPA-secure
// public-key encryption scheme underlying Kyber.
func IndcpaKeypair(paramsK int) ([]byte, []byte, error) {
	skpv := PolyvecNew(paramsK)
	pkpv := PolyvecNew(paramsK)
	e := PolyvecNew(paramsK)
	buf := make([]byte, 2*paramsSymBytes)
	h := sha3.New512()
	_, err := rand.Read(buf[:paramsSymBytes])
	if err != nil {
		return []byte{}, []byte{}, err
	}
	_, err = h.Write(buf[:paramsSymBytes])
	if err != nil {
		return []byte{}, []byte{}, err
	}
	buf = buf[:0]
	buf = h.Sum(buf)
	publicSeed, noiseSeed := buf[:paramsSymBytes], buf[paramsSymBytes:]
	a, err := IndcpaGenMatrix(publicSeed, false, paramsK)
	if err != nil {
		return []byte{}, []byte{}, err
	}
	var nonce byte
	for i := 0; i < paramsK; i++ {
		skpv[i] = polyGetNoise(noiseSeed, nonce, paramsK)
		nonce = nonce + 1
	}
	for i := 0; i < paramsK; i++ {
		e[i] = polyGetNoise(noiseSeed, nonce, paramsK)
		nonce = nonce + 1
	}

	PolyvecNtt(skpv, paramsK)
	PolyvecReduce(skpv, paramsK)
	PolyvecNtt(e, paramsK)
	for i := 0; i < paramsK; i++ {
		pkpv[i] = polyToMont(PolyvecPointWiseAccMontgomery(a[i], skpv, paramsK))
	}

	polyvecAdd(pkpv, e, paramsK)
	PolyvecReduce(pkpv, paramsK)
	return IndcpaPackPrivateKey(skpv, paramsK), IndcpaPackPublicKey(pkpv, publicSeed, paramsK), nil
}

func WritePolyVec(s PolyVec, f *os.File) {
	for _, poly := range s {
		for i, coef := range poly {
			if i > 255 {
				break
			}
			f.WriteString(fmt.Sprintf("%d ", coef))
		}
		f.WriteString(fmt.Sprintf("\n"))
	}
}

// IndcpaEncrypt is the encryption function of the CPA-secure
// public-key encryption scheme underlying Kyber.
func IndcpaEncrypt(m []byte, publicKey []byte, coins []byte, paramsK int) ([]byte, error) {
	sp := PolyvecNew(paramsK)
	ep := PolyvecNew(paramsK)
	bp := PolyvecNew(paramsK)
	publicKeyPolyvec, seed := IndcpaUnpackPublicKey(publicKey, paramsK)
	k := PolyFromMsg(m)
	at, err := IndcpaGenMatrix(seed[:paramsSymBytes], true, paramsK)
	if err != nil {
		return []byte{}, err
	}

	for i := 0; i < paramsK; i++ {
		sp[i] = polyGetNoise(coins, byte(i), paramsK)
		ep[i] = polyGetNoise(coins, byte(i+paramsK), 3)
	}
	epp := polyGetNoise(coins, byte(paramsK*2), 3)
	PolyvecNtt(sp, paramsK)
	PolyvecReduce(sp, paramsK)
	for i := 0; i < paramsK; i++ {
		bp[i] = PolyvecPointWiseAccMontgomery(at[i], sp, paramsK)
	}

	v := PolyvecPointWiseAccMontgomery(publicKeyPolyvec, sp, paramsK)
	PolyvecInvNttToMont(bp, paramsK)
	v = PolyInvNttToMont(v)
	polyvecAdd(bp, ep, paramsK)
	v = PolyAdd(PolyAdd(v, epp), k)
	PolyvecReduce(bp, paramsK)
	return IndcpaPackCiphertext(bp, PolyReduce(v), paramsK), nil
}

// IndcpaDecrypt is the decryption function of the CPA-secure
// public-key encryption scheme underlying Kyber.
func IndcpaDecrypt(c []byte, privateKey []byte, paramsK int) []byte {
	bp, v := IndcpaUnpackCiphertext(c, paramsK)
	privateKeyPolyvec := IndcpaUnpackPrivateKey(privateKey, paramsK)
	PolyvecNtt(bp, paramsK)
	mp := PolyvecPointWiseAccMontgomery(privateKeyPolyvec, bp, paramsK)
	mp = PolyInvNttToMont(mp)
	mp = PolySub(v, mp)
	mp = PolyReduce(mp)
	return PolyToMsg(mp)
}
