package signature

import (
	"CS/Cryptography/rsa"
	"crypto/sha256"
)

func signMessage(plaintext string, publicKey rsa.Key) (signature []uint64) {
	hash := sha256.Sum256([]byte(plaintext))
	stringHash := string(hash[:])
	binaryHash := rsa.SplitString(stringHash)
	signature = rsa.RSAEncryption(binaryHash, publicKey)
	return
}

func verifySignature(message string, signature []uint64, privateKey rsa.Key) bool {
	signatureHash := rsa.MergeToString(rsa.RSAEncryption(signature, privateKey))
	messageHash := sha256.Sum256([]byte(message))
	return signatureHash == string(messageHash[:])
}
