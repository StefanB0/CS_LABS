package pkg

func ViginereEncrypt(s, key string) string{
	newS := ""
	ki := 0
	for i, ch := range s {
		if isLetter(ch) {
			var base, kbase byte
			if isUppercase(ch) {
				base = 'A'
			} else {
				base = 'a'
			}

			if isUppercase(rune(key[ki % len(key)])) {
				kbase = 'A'
			} else {
				kbase = 'a'
			}

			ech := ((s[i] - base + key[ki % len(key)] - kbase) % 26) + base
			ki++
			newS += string(ech)
		} else {
			newS += string(ch)
		}
	}

	return newS
}

func ViginereDecrypt(s, key string) string {
	newS := ""
	ki := 0
	for i, ch := range s {
		if isLetter(ch) {
			var base, kbase byte
			if isUppercase(ch) {
				base = 'A'
			} else {
				base = 'a'
			}

			if isUppercase(rune(key[ki % len(key)])) {
				kbase = 'A'
			} else {
				kbase = 'a'
			}

			dch := ((s[i] - base - (key[ki % len(key)] - kbase) + 26) % 26) + base
			ki++
			newS += string(dch)
		} else {
			newS += string(ch)
		}
	}

	return newS
}
