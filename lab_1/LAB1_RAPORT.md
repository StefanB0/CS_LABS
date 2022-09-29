# Classical Ciphers

### Course: Cryptography & Security
### Author: Boicu Stefan

----

## Objectives:

* Learn the theory and principles of classical ciphers
* Implement 4 classical ciphers


## Implementation description
Caesar Cipher is the simplest form of encryption, just shift the alphabet left or right. The implementation is straightforward, with some boilerplate for string manipulation

### Key settings

First, install the golang language on your system. [Link for download](https://go.dev/learn/)

To run the code, `$cd` into the directory containing `main.go`. Then run the following command.

`$go run main.go` 

This laboratory work also has test coverage. To verify that all the tests run correctly run. Tests account for some edgecases

`$go test ./pkg/ -v`

### Cipher logic

Caesar Cipher is the simplest form of encryption, just shift the alphabet left or right. The implementation is straightforward, with some boilerplate for string manipulation

```go
    newMessage += string((r - base + rune(key)) % 26 + base)

```

Caesar Cipher with permutation is much of the same, but on a random permutation of the alphabet. The alphabet permutation is created pseudo-randomly with an integer as the seed. Seed can be used the same as the key for simplicity. During decryption the order of operations is reversed, including the application of the shift and conversion to and from the alphabet permutation. For decryption a reversed alphabet permutation is created

```go
//op is short for operation
//Encryption
    op:= ch - base
    op = alph[op]
    op += rune(key)
    op %= rune(alphSize)
    op += base
    newS += string(op)
//... 
//Decryption
    op := ch - base 
    op += rune(key)
    op %= rune(alphSize)
    op = reverseAlph[op]
    op += base

    newS += string(op)

```

Viginere Cipher is like Caesar Cipher but squared. Each letter is shifted a different amount depending on the passphrase. The encryption and decription can be easily done with modular arithmetics and accounting for ASCII padding.

```go
//Encryption
    ech := ((s[i] - base + key[ki % len(key)] - kbase) % 26) + base
//Decryption
    dch := ((s[i] - base - (key[ki % len(key)] - kbase) + 26) % 26) + base
```

Play Fair cipher is centered around an encryption square which I implemented as a string of length 25. The play fair rules I use are converting J to I and padding with X. Same row and collumn I shift to the right and down respectively.

Creating the Play Fair square (or key).

```go
    key = addNonRepeatingString(key, phrase)
	key = addNonRepeatingString(key, defaultAlphabet)
```

The most important part of the algorithm is the Play Fair switch function which encrypts pairs of characters. This is the only one of three cyphers that accepts exclusively caps letters (plaintext is converted to be suitable automatically)

p1 is the first letter of the pair and p2 is the second letter of the pair. shift is +1 during encryption and -1 during decryption

```go
switch {
	case p1.x == p2.x:
		p1.y, p2.y = p1.y+shift, p2.y+shift
	case p1.y == p2.y:
		p1.x, p2.x = p1.x+shift, p2.x+shift
	default:
		p1.x, p2.x = p2.x, p1.x
}
```

The algorithm itself simply calls Switch for every pair of letters

```go
for i := 0; i < len(message); i+=2 {
    newMessage += fairSwitch(message[i:i+2], key, decrypt)
}
```



## Conclusions / Screenshots / Results

As part of this laboratory work I learned about classical ciphers and how to implement them. I also learned to do unit tests