# Classical Ciphers

### Course: Cryptography & Security
### Author: Boicu Stefan

----

## Objectives:

* Learn the theory and principles of stream ciphers and classical ciphers
* Implement a stream cipher and a block cipher of my choice

### Key settings

First, install the golang language on your system. [Link for download](https://go.dev/learn/)

This laboratory work also has test coverage. Tests prove that the ciphers work as intended. To verify that all the tests run correctly run.

`$go test ./pkg/ -v`

# RC4

## Theory

RC4 is a stream cipher used to encrypt a stream of bytes of arbitrary length.

RC4 produces a pseudorandom stream of bits based on the provided key and then perfoms a bitwise OR operation on each bit of the plaintext. As a result we get the cipher text. Decription is done by performing the same operation again.

The pseudorandom stream of bits used for encoding is made by first using the Key-Scheduling algorithm (KSA) on the provided key. This mixes the key and results in the initial permutation of bits of length 256. After that the resulting array is fed into the Pseudo random generation algorithm (PRGA) to produce as many bits as needed to encrypt the plaintext.

## Implementation description

First I make and populate an initial slice s0 with the following function. I'll need it later for PRGA

```go
func initialSlice() []byte {
	s := make([]byte, 256)
	for i := 0; i < 256; i++ {
		s[i] = byte(i)
	}
	return s
}
```

Then I permute the secret key provided to create a key array of length 256

```go
func secretKeyArray(key []byte) []byte {
	s := make([]byte, 256)
	for i := 0; i < len(s); i++ {
		s[i] = key[i % len(key)]
	}
	return s
}
```

Then I perform KSA on the array from previous step

```go
func ksaPermutation(s0, keySlice []byte) []byte {
	j:= 0
	for i := 0; i < len(s0); i++ {
		j = (j +int(s0[i]) + int(keySlice[i])) % 256
		s0[i], s0[j] = s0[j], s0[i]
	}
	return s0
}
```

Last for each bit of the plaintext I generate a new pseudo-random bit using PRGA and perform the bitwise OR operation to encrypt it

```go
func prgaPermutation(streamSlice []byte, i,j *int) byte {
	*i  = (*i+1) % 256
	*j = (*j + int(streamSlice[*i])) % 256
	streamSlice[*i], streamSlice[*j] = streamSlice[*j], streamSlice[*i]
	t := int(streamSlice[*i] + streamSlice[*j]) % 256
	b:= streamSlice[t]
	return b
}
```

```go
    var i, j int
	for k := 0; k < len(plaintext); k++ {
		b := prgaPermutation(streamSlice, &i, &j)
		s[k] = b ^ plaintext[k]
	}
	return s
```

The full process

```go
func RC4Encrypt(plaintext, key []byte) []byte {
	s := make([]byte, len(plaintext))
	s0 := initialSlice()
	keySlice := secretKeyArray(key)
	streamSlice := ksaPermutation(s0, keySlice)

	var i, j int
	for k := 0; k < len(plaintext); k++ {
		b := prgaPermutation(streamSlice, &i, &j)
		s[k] = b ^ plaintext[k]
	}
	return s
}
```

# DES

## Theory

DES is a block cipher which uses a lot of confusing techniques to encrypt a block of 64 bits of data. It is not secure though because it could be cracked by bruteforcing all the keys a couple decades ago. Triple DES is theoretically secure enough.

DES uses both substitution and confusion in order to encrypt the message. It uses a lot of pre-determined tables to permute the contents at different stages. And also uses S-Boxes to substitute the already scrambled bits of the data block.

Decryption is done by applying the subkeys in reverse.

[Better explanation](https://page.math.tu-berlin.de/~kant/teaching/hess/krypto-ws2006/des.htm)

## Implementation

First of all I generate 16 subkeys by using a 64 bit encryption key provided by the user. Only 56 bits are actually used though. If I am decrypting, I also reverse the array of subkeys to be used later.

```go
func generateSubkeys(key uint64, pc1, pc2, ksR []uint8) (subkeys []uint64) {
	subkeys = make([]uint64, 16)
	k0 := permuteBlockUniversal(key, pc1[:], 64)
	c, d := splitNr(k0<<4, 64)
	d = d >> 4
	for i := 0; i < 16; i++ {
		c = (c<<(4+ksR[i]))>>4 | (c >> (28 - ksR[i]))
		d = (d<<(4+ksR[i]))>>4 | (d >> (28 - ksR[i]))
		tempJoin := uint64(c)<<28 | uint64(d)
		subkeys[i] = permuteBlockUniversal(tempJoin, pc2[:], 56)
	}

	return subkeys
}
```


```go
func desEncryption(plaintext uint64, key uint64, encrypt bool) uint64 {
	subkeys := generateSubkeys(key, permutationChoice1[:], permutationChoice2[:], keyShiftRotations[:])
	if !encrypt {
		for i, j := 0, len(subkeys)-1; i < j; i, j = i+1, j-1 {
			subkeys[i], subkeys[j] = subkeys[j], subkeys[i]
		}
	}
    ...

}
```

I use a function to permute the bits of a 64bit number and then shift the result accordingly depending where I use it

```go
func permuteBlockUniversal(src uint64, permutation []uint8, srcSize int) (block uint64) {
	for pos, n := range permutation {
		bit := (src >> (uint8(srcSize) - n)) & 1
		block |= bit
		if pos < len(permutation)-1 {
			block = block << 1
		}
	}
	return
}
```

Then I perform the initial permutation of the plaintext using the initial permutation table.

```go
func desEncryption(plaintext uint64, key uint64, encrypt bool) uint64 {
    ...
	ip := permuteBlockUniversal(plaintext, initialPermutationTable[:], 64)
    ...
}
```

Then I slipt the result in two and iterate over it 16 times. Using the subkeys and plugging it in the f-function. In the f-function I expand the half-blocks to 48 bits and then use them in the S-boxes to encrypt it all

```go
func sBoxPerm(eRK uint64, sBoxes [8][4][16]byte) (sb uint32) {
	eRK_Groups := make([]byte, 8)
	sb = 0

	for i := 0; i < 8; i++ {
		eRK_Groups[i] = byte((eRK << (16 + i*6)) >> 58)

		b00 := eRK_Groups[i]
		bx := (b00 << 3) >> 4
		by := ((b00 & 0b00100000) >> 4) | (b00 & 1)
		b1 := sBoxes[i][by][bx]

		sb = sb | uint32(b1)<<((8-i-1)*4)
	}

	return
}
```

```go
func fFunction(r uint32, k uint64, expP, p_table []byte) (fblock uint32) {
	er := permuteBlockUniversal(uint64(r), expP[:], 32)
	eRK := er ^ k
	sp := sBoxPerm(eRK, sBox)
	fblock = uint32(permuteBlockUniversal(uint64(sp), p_table[:], 32))
	return
}
```

```go
func des_Iteration(l0, r0 uint32, subkey uint64, exp_table, p_table []byte) (l_iteration uint32, r_iteration uint32) {
	l_iteration = r0
	r_iteration = l0 ^ fFunction(r0, subkey, exp_table, p_table)
	return
}
```

In this step the sybkeys are inversed if we are decrypting

```go
func desEncryption(plaintext uint64, key uint64, encrypt bool) uint64 {
    ...
	left, right := splitNr(ip, 64)
	leftIterations[0], rightIterations[0] = des_Iteration(left, right, subkeys[0], expansionTable[:], p_table_test[:])
	for i := 1; i < 16; i++ {
		leftIterations[i], rightIterations[i] = des_Iteration(leftIterations[i-1], rightIterations[i-1], subkeys[i], expansionTable[:], p_table_test[:])
	}
    ...
}
```

At long last I reverse the order of the half-blocks which I got from the previous steps, put them in the final permutation table and here we have our cipher text

```go
func desEncryption(plaintext uint64, key uint64, encrypt bool) uint64 {
    ...
    r16l16 := mergeNr(uint64(rightIterations[15]), uint64(leftIterations[15]), 64)
	finalMessage := permuteBlockUniversal(r16l16, finalPermutationTable[:], 64)
    ...
}
```
## Conclusions

As part of this laboratory work I learned about stream and block ciphers and how to implement them.