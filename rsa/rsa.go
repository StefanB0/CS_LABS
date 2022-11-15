package rsa

type Key struct {
	ed uint64
	n  uint64
}

type PrimePair struct {
	x uint64
	y uint64
}

func RSAEncryption(plaintext []uint64, rsaKey Key) []uint64 {
	ciphertext := make([]uint64, len(plaintext))
	for i, block := range plaintext {
		cipherBlock := moduloPow(block, rsaKey.ed, rsaKey.n)
		ciphertext[i] = cipherBlock
	}
	return ciphertext
}

func GenerateKeyPair(primeA, primeB uint64) (privateKey Key, publicKey Key) {
	n := primeA * primeB
	phi := (primeA - 1) * (primeB - 1)
	e := findCoprime(phi)
	var x, y int64
	exgcd(int64(e),int64(phi), &x, &y)
	d := (uint64(x) + phi) % phi

	publicKey = Key{e, n}
	privateKey = Key{d, n}

	return privateKey, publicKey
}

func findCoprime(phi uint64) uint64 {
	for i := uint64(3); i < phi; i++ {
		if gcd(i, phi) == 1 {
			return i
		}
	}
	return phi - 1
}

func gcd(a, b uint64) uint64 {
	for {
		temp := a % b
		if temp == 0 {
			return b
		}
		a = b
		b = temp
	}
}

func exgcd(a, b int64, x, y *int64) int64 {
	// base case
	if b == 0 {
		*x = 1
		*y = 0
		return a
	}

	var x1 int64 = 1
	var y1 int64 = 1
	d := exgcd(b, a%b, &x1, &y1)
	*x = y1
	*y = x1 - y1*(a/b)
	return d
}

func moduloPow(a, b, mod uint64) uint64 {
	var result uint64 = 1
	for i := uint64(0); i < b; i++ {
		result = (result * a) % mod
	}
	return result
}

func SplitString(plaintext string) []uint64 {
	binaryText := make([]uint64, len(plaintext))
	byteText := []byte(plaintext)
	for i := 0; i < len(plaintext); i++ {
		binaryText[i] = uint64(byteText[i])
	}
	return binaryText
}

func MergeToString(binaryMessage []uint64) string {
	byteMessage := make([]byte, len(binaryMessage))
	for i, c := range binaryMessage {
		byteMessage[i] = byte(c)
	}
	return string(byteMessage)
}
