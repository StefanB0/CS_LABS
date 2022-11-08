package main

import (
	"CS/Lab_3/rsa"
	"fmt"
)

func main() {
	p := []int64{7, 13, 3, 35}
	q := []int64{11, 17, 13, 15}

	for i := 0; i < len(p); i++ {
		n := p[i] * q[i]
		phi := (p[i] - 1) * (q[i] - 1)
		e := rsa.FindCoprime(phi)
		var x1, y1 int64
		rsa.Exgcd(e, phi, &x1, &y1)
		d := (x1 + phi) % phi

		fmt.Printf("Key parameters: p: %v, q: %v, n: %v, phi: %v, e: %v, d: %v, x: %v, y: %v\n\n",
			p[i], q[i], n, phi, e, d, x1, y1)
	}

	//TODO Make the strings encode letter by letter, and limit RSA to 8 bits at a time
}
