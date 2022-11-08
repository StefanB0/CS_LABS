package rsa

type Key struct {
	ed int64
	n  int64
}

type PrimePair struct {
	x int64
	y int64
}

func RSAEncryption(plaintext []int64, rsaKey Key) []int64 {
	ciphertext := make([]int64, len(plaintext))
	for i, block := range plaintext {
		cipherBlock := moduloPow(int64(block), rsaKey.ed, rsaKey.n)
		ciphertext[i] = cipherBlock
	}
	return ciphertext
}

func generateKeyPair(primeA, primeB int64) (privateKey Key, publicKey Key) {
	n := primeA * primeB
	phi := (primeA - 1) * (primeB - 1)
	e := FindCoprime(phi)
	var x, y int64
	Exgcd(e, phi, &x, &y)
	d := (x + phi) % phi

	publicKey = Key{e, n}
	privateKey = Key{d, n}

	return privateKey, publicKey
}

func FindCoprime(phi int64) int64 {
	for i := int64(3); i < phi; i++ {
		if Gcd(i, phi) == 1 {
			return i
		}
	}
	return phi - 1
}

func Gcd(a, b int64) int64 {
	for {
		temp := a % b
		if temp == 0 {
			return b
		}
		a = b
		b = temp
	}
}

func Exgcd(a, b int64, x, y *int64) int64 {
	// base case
	if b == 0 {
		*x = 1
		*y = 0
		return a
	}

	var x1 int64 = 1
	var y1 int64 = 1
	d := Exgcd(b, a%b, &x1, &y1)
	*x = y1
	*y = x1 - y1*(a/b)
	return d
}

func moduloPow(a, b, mod int64) int64 {
	var result int64 = 1
	for i := int64(0); i < b; i++ {
		result = (result * a) % mod
	}
	return result
}

func splitString(plaintext string) []int64 {
	binaryText := make([]int64, len(plaintext))
	byteText := []byte(plaintext)
	for i := 0; i < len(plaintext); i++ {
		binaryText[i] = int64(byteText[i])
	}
	return binaryText
}

func mergeToString(binaryMessage []int64) string {
	byteMessage := make([]byte, len(binaryMessage))
	for i, c := range binaryMessage {
		byteMessage[i] = byte(c)
	}
	return string(byteMessage)
}
