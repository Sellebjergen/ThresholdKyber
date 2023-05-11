package kyberk2so

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/sha3"
)

// IndcpaKeypair generates public and private keys for the CPA-secure
// public-key encryption scheme underlying Kyber.
func IndcpaKeypair_nocomp(paramsK int) (PolyVec, PolyVec, []byte) {
	skpv := PolyvecNew(paramsK)
	pkpv := PolyvecNew(paramsK)
	e := PolyvecNew(paramsK)
	buf := make([]byte, 2*paramsSymBytes)
	h := sha3.New512()
	_, _ = rand.Read(buf[:paramsSymBytes])
	_, _ = h.Write(buf[:paramsSymBytes])
	buf = buf[:0]
	buf = h.Sum(buf)
	publicSeed, noiseSeed := buf[:paramsSymBytes], buf[paramsSymBytes:]
	a, _ := IndcpaGenMatrix(publicSeed, false, paramsK)
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
	return skpv, pkpv, publicSeed
}

// IndcpaEncrypt is the encryption function of the CPA-secure
// public-key encryption scheme underlying Kyber.
func IndcpaEncrypt_nocomp(m []byte, publicKey PolyVec, seed []byte, coins []byte, paramsK int) (PolyVec, Poly) {
	sp := PolyvecNew(paramsK)
	ep := PolyvecNew(paramsK)
	bp := PolyvecNew(paramsK)
	k := PolyFromMsg(m)
	fmt.Println("k")
	fmt.Println(k)
	at, _ := IndcpaGenMatrix(seed[:paramsSymBytes], true, paramsK)

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

	v := PolyvecPointWiseAccMontgomery(publicKey, sp, paramsK)
	PolyvecInvNttToMont(bp, paramsK)
	v = PolyInvNttToMont(v)
	polyvecAdd(bp, ep, paramsK)
	v = PolyAdd(PolyAdd(v, epp), k)
	PolyvecReduce(bp, paramsK)
	return bp, PolyReduce(v)
}

// IndcpaDecrypt is the decryption function of the CPA-secure
// public-key encryption scheme underlying Kyber.
func IndcpaDecrypt_nocomp(u PolyVec, v Poly, privateKey PolyVec, paramsK int) []byte {
	PolyvecNtt(u, paramsK)
	mp := PolyvecPointWiseAccMontgomery(privateKey, u, paramsK)
	mp = PolyInvNttToMont(mp)
	mp = PolySub(v, mp)
	mp = PolyReduce(mp)
	fmt.Println("mp")
	fmt.Println(mp)
	return PolyToMsg(mp)
}
