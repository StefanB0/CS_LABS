package pkg

import "strings"

func PlayFairCypher(message, phrase string, decrypt bool) string {
	key := PlayFairKey(phrase)
	message = adjustMessage(message)
	newMessage := ""

	for i := 0; i < len(message); i+=2 {
		newMessage += fairSwitch(message[i:i+2], key, decrypt)
	}
	return newMessage
}

func fairSwitch(pair, key string, decrypt bool) string {
	var p1, p2 struct{ xy, x, y int }
	var newPair string
	
	shift := 1
	if decrypt {
		shift = -1
	}

	p1.xy = strings.IndexByte(key, pair[0])
	p2.xy = strings.IndexByte(key, pair[1])
	p1.x, p1.y = p1.xy%5, p1.xy/5
	p2.x, p2.y = p2.xy%5, p2.xy/5

	switch {
	case p1.x == p2.x:
		p1.y, p2.y = p1.y+shift, p2.y+shift
	case p1.y == p2.y:
		p1.x, p2.x = p1.x+shift, p2.x+shift
	default:
		p1.x, p2.x = p2.x, p1.x
	}
	
	bundle := []*int{&p1.x, &p1.y, &p2.x, &p2.y}
	for _, coor := range bundle {
		if *coor >= 5 {
			*coor = 0
		}
		
		if *coor < 0 {
			*coor = 4
		}
	}

	newPair = string(key[p1.x+p1.y*5]) + string(key[p2.x+p2.y*5])
	return newPair
}

func adjustMessage(message string) string {
	message0 := strings.Replace(message, "J", "I", -1)
	message1 := processPlaintext(message0)

	l := len(message1)
	for i := 0; i < l-1; i++ {
		if message1[i] == message1[i+1] {
			l++
			message1 = instertString(message1, 'X', i+1)
		}
		i++
	}

	if len(message1)%2 != 0 {
		message1 += "X"
	}

	return message1
}

func PlayFairKey(phrase string) string {
	key := "J"
	defaultAlphabet := "abcdefghijklmnopqrstuvwxyz"
	phrase = processPlaintext(phrase)
	phrase = strings.Replace(phrase, "J", "I", -1)

	key = addNonRepeatingString(key, phrase)
	key = addNonRepeatingString(key, defaultAlphabet)
	key = key[1:]

	return key
}

func addNonRepeatingString(os, s string) string {
	s = processPlaintext(s)
	for _, r := range s {
		if !strings.ContainsRune(os, uppercase(r)) {
			os = os + string(uppercase(r))
		}
	}
	return os
}
