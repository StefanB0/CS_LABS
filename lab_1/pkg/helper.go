package pkg

func processPlaintext(s string) string {
	newS := ""

	for _, ch := range s {
		if isLetter(ch) {
			newS += string(uppercase(ch))
		}
	}
	return newS
}

func uppercase(ch rune) rune {
	if ch >= 97 && ch <= 122 {
		return ch - 32
	}

	return ch
}

func isLetter(r rune) bool {
	if (r >= 97 && r <= 122) || (r >= 65 && r <= 90) {
		return true
	} else {
		return false
	}
}

func isUppercase(r rune) bool {
	if r >= 65 && r <= 90 {
		return true
	} else {
		return false
	}
}

func instertString(s string, ch byte, i int) string {
	newS := s[:i]
	newS += string(ch)
	newS += s[i:]

	return newS
}