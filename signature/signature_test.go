package signature

import (
	"CS/Cryptography/rsa"
	"testing"
)

func TestSignature(t *testing.T) {
	privatekey, publicKey := rsa.GenerateKeyPair(1087, 563)
	tables := []struct {
		message string
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
