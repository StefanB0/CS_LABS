package main

import (
	"CS/Lab_1/pkg"
	"fmt"
)

func main() {
	message:= "Hello world"
	keyInt := 3
	passphrase:= "FAF"
	fmt.Println("Message:",  message, "secret: FAF or 3")
	fmt.Println("Caesar: ", pkg.CaesarEncrypt(message, keyInt))
	fmt.Println("Caesar with permutation:", pkg.CaesarPermEncrypt(message, keyInt, pkg.PermutateAlphabet(keyInt)))
	fmt.Println("Viginere Cipher:", pkg.ViginereEncrypt(message, passphrase))
	fmt.Println("PlayFair Cipher:", pkg.PlayFairCypher(message, passphrase, false))

}