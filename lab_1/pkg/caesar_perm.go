package pkg

import (
	"math/rand"
)

func PermutateAlphabet(key int) []rune {
	newA := []rune{}
	for i:= 0; i < 26; i++ {
		newA = append(newA, rune(i))
	}

	rand.Seed(int64(key))
	for i := 25; i >= 0; i-- {
		j := rand.Int() % 26
		temp := newA[i]
		newA[i] = newA[j]
		newA[j] = temp
	}

	return newA
}

func CaesarPermEncrypt(s string, key int, alph []rune) string {
	alphSize := 26
	key = fixKey(key)
	newS := ""
	for _, ch := range s {
		if isLetter(ch) {
			var base rune
			if isUppercase(ch) {
				base = 'A'
			} else {
				base = 'a'
			}
			op:= ch - base
			op = alph[op]
			op += rune(key)
			op %= rune(alphSize)
			op += base
			newS += string(op)
		} else {
			newS += string(ch)
		}
	}

	return newS
}

func CaesarPermDecrypt(s string, key int, alph []rune) string {
	alphSize := 26
	key *= -1
	key = fixKey(key)

	var reverseAlph []rune
	for i:= 0; i < alphSize; i++ {
		reverseAlph = append(reverseAlph, rune(i))
	}

	for i := range alph {
		reverseAlph[alph[i]] = rune(i)
	}

	newS := ""
	for _, ch := range s {
		if isLetter(ch) {
			var base rune
			if isUppercase(ch) {
				base = 'A'
			} else {
				base = 'a'
			}

			op := ch - base 
			op += rune(key)
			op %= rune(alphSize)
			op = reverseAlph[op]
			op += base

			newS += string(op)
		} else {
			newS += string(ch)
		}
	}

	return newS
}
