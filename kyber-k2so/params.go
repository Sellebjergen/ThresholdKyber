/* SPDX-FileCopyrightText: © 2020-2021 Nadim Kobeissi <nadim@symbolic.software>
 * SPDX-License-Identifier: MIT */

package kyberk2so

const ParamsN int = 256
const ParamsQ int = 3329
const ParamsK int = 2
const paramsQinv int = 62209
const ParamsPolyBytes int = 384 // 12 * 256 / 8 = 384

const paramsSymBytes int = 32
const paramsETAK512 int = 3
const paramsETAK768K1024 int = 2
const paramsPolyvecBytesK512 int = 2 * ParamsPolyBytes
const paramsPolyvecBytesK768 int = 3 * ParamsPolyBytes
const paramsPolyvecBytesK1024 int = 4 * ParamsPolyBytes
const paramsPolyvecBytesK1280 int = 5 * ParamsPolyBytes
const paramsPolyvecBytesK1536 int = 6 * ParamsPolyBytes
const paramsPolyvecBytesK1792 int = 7 * ParamsPolyBytes
const paramsPolyvecCompressedBytesK512 int = 2 * 320
const paramsPolyvecCompressedBytesK768 int = 3 * 320
const paramsPolyvecCompressedBytesK1024 int = 4 * 352
const paramsPolyvecCompressedBytesK1280 int = 5 * 320 // Change
const paramsPolyvecCompressedBytesK1536 int = 6 * 320 // Change
const paramsPolyvecCompressedBytesK1792 int = 7 * 352 // Change

// herfra bliver det brugt i KEM
const paramsPolyCompressedBytesK512 int = 128
const paramsPolyCompressedBytesK768 int = 128
const paramsPolyCompressedBytesK1024 int = 160

const paramsIndcpaPublicKeyBytesK512 int = paramsPolyvecBytesK512 + paramsSymBytes
const paramsIndcpaPublicKeyBytesK768 int = paramsPolyvecBytesK768 + paramsSymBytes
const paramsIndcpaPublicKeyBytesK1024 int = paramsPolyvecBytesK1024 + paramsSymBytes
const paramsIndcpaSecretKeyBytesK512 int = 2 * ParamsPolyBytes
const paramsIndcpaSecretKeyBytesK768 int = 3 * ParamsPolyBytes
const paramsIndcpaSecretKeyBytesK1024 int = 4 * ParamsPolyBytes

// Kyber512SKBytes is a constant representing the byte length of private keys in Kyber-512.
const Kyber512SKBytes int = paramsPolyvecBytesK512 + ((paramsPolyvecBytesK512 + paramsSymBytes) + 2*paramsSymBytes)

// Kyber768SKBytes is a constant representing the byte length of private keys in Kyber-768.
const Kyber768SKBytes int = paramsPolyvecBytesK768 + ((paramsPolyvecBytesK768 + paramsSymBytes) + 2*paramsSymBytes)

// Kyber1024SKBytes is a constant representing the byte length of private keys in Kyber-1024.
const Kyber1024SKBytes int = paramsPolyvecBytesK1024 + ((paramsPolyvecBytesK1024 + paramsSymBytes) + 2*paramsSymBytes)

// Kyber512PKBytes is a constant representing the byte length of public keys in Kyber-512.
const Kyber512PKBytes int = paramsPolyvecBytesK512 + paramsSymBytes

// Kyber768PKBytes is a constant representing the byte length of public keys in Kyber-768.
const Kyber768PKBytes int = paramsPolyvecBytesK768 + paramsSymBytes

// Kyber1024PKBytes is a constant representing the byte length of public keys in Kyber-1024.
const Kyber1024PKBytes int = paramsPolyvecBytesK1024 + paramsSymBytes

// Kyber512CTBytes is a constant representing the byte length of ciphertexts in Kyber-512.
const Kyber512CTBytes int = paramsPolyvecCompressedBytesK512 + paramsPolyCompressedBytesK512

// Kyber768CTBytes is a constant representing the byte length of ciphertexts in Kyber-768.
const Kyber768CTBytes int = paramsPolyvecCompressedBytesK768 + paramsPolyCompressedBytesK768

// Kyber1024CTBytes is a constant representing the byte length of ciphertexts in Kyber-1024.
const Kyber1024CTBytes int = paramsPolyvecCompressedBytesK1024 + paramsPolyCompressedBytesK1024

// KyberSSBytes is a constant representing the byte length of shared secrets in Kyber.
const KyberSSBytes int = 32
