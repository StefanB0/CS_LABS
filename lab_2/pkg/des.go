package pkg

const (
	BLOCKSIZE = 64
)

func DesEncrypt(plaintext uint64, key uint64) uint64 {
	return desEncryption(plaintext, key, false)
}

func DesDecrypt(plaintext uint64, key uint64) uint64 {
	return desEncryption(plaintext, key, true)
}

func desEncryption(plaintext uint64, key uint64, encrypt bool) uint64 {
	subkeys := generateSubkeys(key, permutationChoice1[:], permutationChoice2[:], keyShiftRotations[:])
	if !encrypt {
		for i, j := 0, len(subkeys)-1; i < j; i, j = i+1, j-1 {
			subkeys[i], subkeys[j] = subkeys[j], subkeys[i]
		}
	}

	ip := permuteBlockUniversal(plaintext, initialPermutationTable[:], 64)
	left, right := splitNr(ip, 64)

	leftIterations := make([]uint32, 16)
	rightIterations := make([]uint32, 16)

	leftIterations[0], rightIterations[0] = des_Iteration(left, right, subkeys[0], expansionTable[:], p_table_test[:])
	for i := 1; i < 16; i++ {
		leftIterations[i], rightIterations[i] = des_Iteration(leftIterations[i-1], rightIterations[i-1], subkeys[i], expansionTable[:], p_table_test[:])
	}

	r16l16 := mergeNr(uint64(rightIterations[15]), uint64(leftIterations[15]), 64)

	finalMessage := permuteBlockUniversal(r16l16, finalPermutationTable[:], 64)
	return finalMessage
}

func splitNr(message uint64, size int) (uint32, uint32) {
	left := uint32(message >> (size / 2))
	right := uint32(message)
	return left, right
}

func mergeNr(left, right uint64, size int) uint64 {
	return left<<(size/2) | right
}

func generateSubkeys(key uint64, pc1, pc2, ksR []uint8) (subkeys []uint64) {
	subkeys = make([]uint64, 16)
	k0 := permuteBlockUniversal(key, pc1[:], 64)
	c, d := splitNr(k0<<4, 64)
	d = d >> 4
	for i := 0; i < 16; i++ {
		c = (c<<(4+ksR[i]))>>4 | (c >> (28 - ksR[i]))
		d = (d<<(4+ksR[i]))>>4 | (d >> (28 - ksR[i]))
		tempJoin := uint64(c)<<28 | uint64(d)
		subkeys[i] = permuteBlockUniversal(tempJoin, pc2[:], 56)
	}

	return subkeys
}

func permuteBlockUniversal(src uint64, permutation []uint8, srcSize int) (block uint64) {
	for pos, n := range permutation {
		bit := (src >> (uint8(srcSize) - n)) & 1
		block |= bit
		if pos < len(permutation)-1 {
			block = block << 1
		}
	}
	return
}

func des_Iteration(l0, r0 uint32, subkey uint64, exp_table, p_table []byte) (l_iteration uint32, r_iteration uint32) {
	l_iteration = r0
	r_iteration = l0 ^ fFunction(r0, subkey, exp_table, p_table)
	return
}

func fFunction(r uint32, k uint64, expP, p_table []byte) (fblock uint32) {
	er := permuteBlockUniversal(uint64(r), expP[:], 32)
	eRK := er ^ k
	sp := sBoxPerm(eRK, sBox)
	fblock = uint32(permuteBlockUniversal(uint64(sp), p_table[:], 32))
	return
}

func sBoxPerm(eRK uint64, sBoxes [8][4][16]byte) (sb uint32) {
	eRK_Groups := make([]byte, 8)
	sb = 0

	for i := 0; i < 8; i++ {
		eRK_Groups[i] = byte((eRK << (16 + i*6)) >> 58)

		b00 := eRK_Groups[i]
		bx := (b00 << 3) >> 4
		by := ((b00 & 0b00100000) >> 4) | (b00 & 1)
		b1 := sBoxes[i][by][bx]

		sb = sb | uint32(b1)<<((8-i-1)*4)
	}

	return
}

func appendBytes(plaintext *[]uint64) {
	tail := 0
	l := len(*plaintext)
	if l%BLOCKSIZE != 0 {
		tail = (BLOCKSIZE - l) % BLOCKSIZE
	} else {
		return
	}
	tailspace := make([]uint64, tail)
	*plaintext = append(*plaintext, tailspace...)
}

func splitStringBlocks(plaintext string) (messageA []uint64) {
	for len(plaintext)%8 != 0 {
		plaintext += "0"
	}

	messageA = make([]uint64, len(plaintext)/8)

	for i := 0; i < len(plaintext)/8; i++ {
		var block uint64
		for j := 0; j < 8; j++ {
			block |= uint64([]byte(plaintext)[i*8+j]) << ((7 - j) * 8)
		}
		messageA[i] = block
	}
	return
}

func mergeStringBlocks(messageA []uint64) (messageS string) {
	messageS = ""
	for i := 0; i < len(messageA); i++ {
		var stringBlock string
		for j := 0; j < 8; j++ {
			stringBlock += string(byte(messageA[i] >> ((7 - j) * 8)))
		}
		messageS += stringBlock
	}
	return
}
