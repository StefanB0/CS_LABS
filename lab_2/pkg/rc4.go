package pkg

func initialSlice() []byte {
	s := make([]byte, 256)
	for i := 0; i < 256; i++ {
		s[i] = byte(i)
	}
	return s
}

func secretKeyArray(key []byte) []byte {
	s := make([]byte, 256)
	for i := 0; i < len(s); i++ {
		s[i] = key[i % len(key)]
	}
	return s
}

func ksaPermutation(s0, keySlice []byte) []byte {
	j:= 0
	for i := 0; i < len(s0); i++ {
		j = (j +int(s0[i]) + int(keySlice[i])) % 256
		s0[i], s0[j] = s0[j], s0[i]
	}
	return s0
}

func prgaPermutation(streamSlice []byte, i,j *int) byte {
	*i  = (*i+1) % 256
	*j = (*j + int(streamSlice[*i])) % 256
	streamSlice[*i], streamSlice[*j] = streamSlice[*j], streamSlice[*i]
	t := int(streamSlice[*i] + streamSlice[*j]) % 256
	b:= streamSlice[t]
	return b
}

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

func RC4Decrypt(ciphertext, key []byte) []byte {
	return RC4Encrypt(ciphertext, key)
}
