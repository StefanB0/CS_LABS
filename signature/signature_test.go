package signature

import (
	"CS/Cryptography/rsa"
	"crypto/sha256"
	"testing"
)

func TestSignature(t *testing.T) {
	privatekey, publicKey := rsa.GenerateKeyPair(1087, 563)
	tables := []struct {
		message  string
	}{
		{"Luke, I am your father"},
		{"Roses are red, violets are blue"},
		{"Why did the chicken cross the road"},
		{"There are two hard things in computer science, cache invalidation and naming things"},
	}
	for _, table := range tables {
		signature := signMessage(table.message, publicKey)
		result := verifySignature(table.message, signature, privatekey)
		if !result {
			t.Errorf("\nSomething went wrong")
		}
	}
}

func TestHashTransform(t *testing.T) {
	tables := []struct {
		message  string
	}{
		{"Luke, I am your father"},
		{"Roses are red, violets are blue"},
		{"Why did the chicken cross the road"},
		{"There are two hard things in computer science, cache invalidation and naming things"},
	}
	for _, table := range tables {
		hash := sha256.Sum256([]byte(table.message))
		transform := hashTransform(hash)
		result := hashDecrypt(transform)
		
		if string(hash[:]) != string(result[:]) {
			t.Errorf("\nHashes differ\n\tinitialHash:\t\t%0x\n\tIntermediateHash:\t%0x\n\tDecryptedHash:\t\t%0x", hash, transform, result)
		}
	}
}

func TestVerifyMessage(t *testing.T){
	privatekey, publicKey := rsa.GenerateKeyPair(174440041, 3657500101)
	tables := []struct {
		message  string
	}{
		{"Luke, I am your father"},
		// {"Roses are red, violets are blue"},
		// {"Why did the chicken cross the road"},
		// {"There are two hard things in computer science, cache invalidation and naming things"},
	}
	for _, table := range tables {
		hash := sha256.Sum256([]byte(table.message))
		signature := signMessage(table.message, publicKey)
		newHash := hashDecrypt(rsa.RSAEncryption(signature, privatekey))
		result := verifySignature(table.message, signature, privatekey)
		if !result {
			t.Errorf("\nSomething went wrong\n\tHash:\t\t%v\n\tNewHash:\t\t%v\n\tSignature:\t\t%08x",hash, newHash, signature)
		}
	}
}
