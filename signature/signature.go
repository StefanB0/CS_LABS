package signature

import (
	"CS/Cryptography/rsa"
	"crypto/sha256"
)

func signMessage(plaintext string, publicKey rsa.Key) (signature []uint64) {
	hash := sha256.Sum256([]byte(plaintext))
	eHash := rsa.SplitString(string(hash[:]))
	return rsa.RSAEncryption(eHash, publicKey)
}

func verifySignature(message string, signature []uint64, privateKey rsa.Key) bool {
	signatureHash := rsa.MergeToString(rsa.RSAEncryption(signature, privateKey))
	messageHash := sha256.Sum256([]byte(message))
	return signatureHash == string(messageHash[:])
}
