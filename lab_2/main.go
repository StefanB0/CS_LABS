package main

import (
	"fmt"
)

func universalPrintS(data uint64, size, chunk int) string {
	var s string

	switch chunk {
	case 4:
		s = "%04b"
	case 6:
		s = "%06b"
	case 7:
		s = "%07b"
	case 8:
		s = "%08b"
	case 12:
		s = "%012b"
	case 14:
		s = "%014b"
	case 16:
		s = "%016b"
	case 24:
		s = "%024b"
	case 32:
		s = "%032b"
	case 48:
		s = "%048b"
	case 64:
		s = "%064b"
	default:
		s = "%064b"
	}

	// 8 in 16 00 00 00 00 11 00 11 00

	result := ""
	for i := 0; i < size/chunk; i++ {
		result += fmt.Sprintf(s+" ", (data<<(chunk*i+(64-size)))>>(64-chunk))
	}
	result += "\n"
	return result
}

func main() {
	fmt.Println(string(byte(0x62)))
}
