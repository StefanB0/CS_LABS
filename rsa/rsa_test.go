package rsa

import (
	"testing"
)

func TestRSAEncryptDecrypt(t *testing.T) {
	tables := []struct {
		plaintext string
		primePair PrimePair
	}{
		{plaintext: "Hi", primePair: PrimePair{53, 59}},
		{plaintext: "Hello world", primePair: PrimePair{1087, 563}},
		{plaintext: "Asta la vista, baby", primePair: PrimePair{419, 809}},
		{plaintext: "Make moldova great again", primePair: PrimePair{283, 191}},
		{plaintext: "I am your father, Luke", primePair: PrimePair{277, 223}},
	}
	for _, table := range tables {
		privateKey, publicKey := GenerateKeyPair(table.primePair.x, table.primePair.y)
		binaryPlaintext := SplitString(table.plaintext)

		enMessage := RSAEncryption(binaryPlaintext, publicKey)
		deBinaryMessage := RSAEncryption(enMessage, privateKey)

		result := MergeToString(deBinaryMessage)

		if result != table.plaintext {
			t.Errorf("Encryption - Decryption process failed:\nexpected:\t%v\ngot:\t\t%v", table.plaintext, result)
		}
	}
}

func TestRSAEncrypt(t *testing.T) {
	tables := []struct {
		plaintext []uint64
		publicKey Key
		expected  []uint64
	}{
		{plaintext: []uint64{123}, publicKey: Key{ed: 17, n: 3233}, expected: []uint64{855}},
	}
	for _, table := range tables {
		result := RSAEncryption(table.plaintext, table.publicKey)
		if !compareU64Slice(result, table.expected) {
			t.Errorf("Encryption failed:\nexpected:\t%v\ngot:\t\t%v", table.expected, result)
		}
	}
}

func TestGenerateKeyPair(t *testing.T) {
	tables := []struct {
		x                  uint64
		y                  uint64
		expectedPrivateKey Key
		expectedPublicKey  Key
	}{
		{x: 07, y: 11, expectedPrivateKey: Key{n: 77, ed: 43}, expectedPublicKey: Key{n: 77, ed: 7}},
		{x: 13, y: 17, expectedPrivateKey: Key{n: 221, ed: 77}, expectedPublicKey: Key{n: 221, ed: 5}},
		{x: 3, y: 13, expectedPrivateKey: Key{n: 39, ed: 5}, expectedPublicKey: Key{n: 39, ed: 5}},
		{x: 3, y: 11, expectedPrivateKey: Key{n: 33, ed: 7}, expectedPublicKey: Key{n: 33, ed: 3}},
	}
	for _, table := range tables {
		privateKey, publicKey := GenerateKeyPair(table.x, table.y)
		if privateKey.ed != table.expectedPrivateKey.ed || privateKey.n != table.expectedPrivateKey.n {
			t.Errorf("\nWrong private key:\nexpected:\td:%v\tn:%v\ngot:\t\te:%v\tn:%v", table.expectedPrivateKey.ed, table.expectedPrivateKey.n, privateKey.ed, privateKey.n)
		}
		if publicKey.ed != table.expectedPublicKey.ed || publicKey.n != table.expectedPublicKey.n {
			t.Errorf("\nWrong public key:\nexpected:\te:%v\tn:%v\ngot:\t\te:%v\tn:%v", table.expectedPublicKey.ed, table.expectedPublicKey.n, publicKey.ed, publicKey.n)
		}
	}
}

func TestModuloPow(t *testing.T) {
	tables := []struct {
		c        uint64
		e        uint64
		n        uint64
		expected uint64
	}{
		{c: 123, e: 17, n: 3233, expected: 855},
	}
	for _, table := range tables {
		result := moduloPow(table.c, table.e, table.n)
		if result != table.expected {
			t.Errorf("\nModulo Power:\nexpected:\t%v\ngot:\t\t%v", table.expected, result)
		}
	}
}

func TestFindCoPrime(t *testing.T) {
	tables := []struct {
		input uint64
	}{
		{input: 60},
		{input: 192},
		{input: 24},
		{input: 476},
		{input: 30},
	}
	for _, table := range tables {
		result := findCoprime(table.input)
		if gcd(table.input, result) > 1 {
			t.Errorf("Result value is not coprime with input. got %v. GCD = %v", result, gcd(table.input, result))
		}
	}
}

func TestExGCD(t *testing.T) {
	tables := []struct {
		phi int64
		e   int64
	}{
		{phi: 60, e: 7},
		{phi: 192, e: 5},
		{phi: 24, e: 5},
		{phi: 476, e: 3},
	}
	for _, table := range tables {
		var x, y int64
		result := exgcd(table.e, table.phi, &x, &y)
		d := (x + int64(table.phi)) % table.phi
		if result != int64(gcd(uint64(table.e), uint64(table.phi))) {
			t.Errorf("Extended gcd not finding the greatest common divisor:\nexpected:\t%v\ngot:\t\t%v", gcd(uint64(table.e), uint64(table.phi)), result)
		}
		if (d*table.e)%table.phi != 1 {
			t.Errorf("Wrong inverse modulo of e:%v. (e * d) mod phi = %v", table.e, (d*table.e)%table.phi)
		}
	}
}

func TestSplitString(t *testing.T) {
	tables := []struct {
		input    string
		expected []uint64
	}{
		{input: "Hello world", expected: []uint64{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64}},
	}
	for _, table := range tables {
		result := SplitString(table.input)
		if !compareU64Slice(result, table.expected) {
			t.Errorf("Failed conversion from string to byte array:\nexpected:\t%v\ngot:\t\t%v", table.expected, result)
		}
	}
}

func TestMergeToString(t *testing.T) {
	tables := []struct {
		input    []uint64
		expected string
	}{
		{input: []uint64{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64}, expected: "Hello world"},
	}
	for _, table := range tables {
		result := MergeToString(table.input)
		if result != table.expected {
			t.Errorf("Failed conversion from byte array to string:\nexpected:\t%v\ngot:\t\t%v", table.expected, result)
		}
	}
}
