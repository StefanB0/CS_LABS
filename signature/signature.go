package signature

import (
	"CS/Cryptography/rsa"
	"crypto/sha256"
)

func hashTransform(hash [32]byte) []uint64 {
	result := make([]uint64, 4)
	for i := 0; i < 4; i++ {
		for j := 0; j < 8; j++ {
			result[i] |= uint64(hash[i*8+j]) << ((7 - j) * 8)
		}
	}
	return result
}

func hashDecrypt(hash64 []uint64) [32]byte {
	result := [32]byte{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 8; j++ {
			result[i*8+j] = byte(hash64[i] >> ((7 - j) * 8))
		}
	}
	return result
}

func signMessage(plaintext string, publicKey rsa.Key) (signature []uint64) {
	hash := sha256.Sum256([]byte(plaintext))
	hash64 := hashTransform(hash)
	return rsa.RSAEncryption(hash64, publicKey)
}

func verifySignature(message string, signature []uint64, privateKey rsa.Key) bool {
	signatureHash := hashDecrypt(rsa.RSAEncryption(signature, privateKey))
	messageHash := sha256.Sum256([]byte(message))
	return string(signatureHash[:]) == string(messageHash[:])
}
