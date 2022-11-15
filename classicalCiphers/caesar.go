package classical

func fixKey(key int) int {
	if key < 0 {
		key *= -1
		key %= 26
		key = 26 - key
	}
	key %= 26

	return key
}

func CaesarEncrypt(message string, key int) string {
	key = fixKey(key)

	newMessage := ""
	for _, r := range message {
		if isLetter(r) {
			var base rune
			if isUppercase(r) {
				base = 'A'
			} else {
				base = 'a'
			}
			newMessage += string((r - base + rune(key)) % 26 + base)
		} else {
			newMessage += string(r)
		}
	}

	return newMessage
}

func CaesarDecrypt(s string, key int) string {
	key *= -1
	return CaesarEncrypt(s, key)
}
