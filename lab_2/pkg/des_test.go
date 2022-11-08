package pkg

import "testing"

var (
	expectedSubkeys = []uint64{
		0b000110110000001011101111111111000111000001110010,
		0b011110011010111011011001110110111100100111100101,
		0b010101011111110010001010010000101100111110011001,
		0b011100101010110111010110110110110011010100011101,
		0b011111001110110000000111111010110101001110101000,
		0b011000111010010100111110010100000111101100101111,
		0b111011001000010010110111111101100001100010111100,
		0b111101111000101000111010110000010011101111111011,
		0b111000001101101111101011111011011110011110000001,
		0b101100011111001101000111101110100100011001001111,
		0b001000010101111111010011110111101101001110000110,
		0b011101010111000111110101100101000110011111101001,
		0b100101111100010111010001111110101011101001000001,
		0b010111110100001110110111111100101110011100111010,
		0b101111111001000110001101001111010011111100001010,
		0b110010110011110110001011000011100001011111110101,
	}
)

func TestDESEncryptDecryptMessage(t *testing.T) {
	tables := []struct {
		plaintext string
		key       uint64
		expected  string
	}{
		{plaintext: "r", key: 0x133457799BBCDFF1, expected: "r0000000"},
		{plaintext: "Hello Me", key: 0x133457799BBCDFF1, expected: "Hello Me"},
		{plaintext: "Hello World", key: 0x133457799BBCDFF1, expected: "Hello World00000"},
		{plaintext: "HHHHHHHHH", key: 0x133457799BBCDFF1, expected: "HHHHHHHHH0000000"},
		{plaintext: lorem_ipsum, key: 0x133457799BBCDFF1, expected: lorem_ipsum + "000"},
	}
	for _, table := range tables {
		formattedPlaintext := splitStringBlocks(table.plaintext)
		encryptedMessage := make([]uint64, len(formattedPlaintext))
		for i, textBlock := range formattedPlaintext {
			encryptedMessage[i] = DesEncrypt(textBlock, table.key)
		}

		unformattedDecryptedMessage := make([]uint64, len(formattedPlaintext))
		for i, textBlock := range encryptedMessage {
			unformattedDecryptedMessage[i] = DesDecrypt(textBlock, table.key)
		}

		decryptedMessage := mergeStringBlocks(unformattedDecryptedMessage)

		if table.expected != decryptedMessage {
			t.Errorf("Decrypted message differs from plaintext\n%s\n\n%s", table.plaintext, decryptedMessage)
		}
	}
}

func TestPermuteBlock(t *testing.T) {
	tables := []struct {
		input     uint64
		expected  uint64
		permTable []byte
	}{
		{
			0b0001001100110100010101110111100110011011101111001101111111110001,
			0b11110000110011001010101011110101010101100110011110001111,
			permutationChoice1[:],
		},
	}
	for _, table := range tables {
		result := permuteBlockUniversal(table.input, table.permTable, 64)
		if result != table.expected {
			t.Errorf("expected:%056b, got:%056b", table.expected, result)
		}
	}
}

func TestSubkeys(t *testing.T) {
	var key uint64 = 0b0001001100110100010101110111100110011011101111001101111111110001
	E_Subkeys := []uint64{
		0b000110110000001011101111111111000111000001110010,
		0b011110011010111011011001110110111100100111100101,
		0b010101011111110010001010010000101100111110011001,
		0b011100101010110111010110110110110011010100011101,
		0b011111001110110000000111111010110101001110101000,
		0b011000111010010100111110010100000111101100101111,
		0b111011001000010010110111111101100001100010111100,
		0b111101111000101000111010110000010011101111111011,
		0b111000001101101111101011111011011110011110000001,
		0b101100011111001101000111101110100100011001001111,
		0b001000010101111111010011110111101101001110000110,
		0b011101010111000111110101100101000110011111101001,
		0b100101111100010111010001111110101011101001000001,
		0b010111110100001110110111111100101110011100111010,
		0b101111111001000110001101001111010011111100001010,
		0b110010110011110110001011000011100001011111110101,
	}

	tables := []struct {
		key              uint64
		expected_subkeys []uint64
		i                int
	}{
		{key, E_Subkeys, 0},
		{key, E_Subkeys, 1},
		{key, E_Subkeys, 2},
		{key, E_Subkeys, 3},
		{key, E_Subkeys, 4},
		{key, E_Subkeys, 5},
		{key, E_Subkeys, 6},
		{key, E_Subkeys, 7},
		{key, E_Subkeys, 8},
		{key, E_Subkeys, 9},
		{key, E_Subkeys, 10},
		{key, E_Subkeys, 11},
		{key, E_Subkeys, 12},
		{key, E_Subkeys, 13},
		{key, E_Subkeys, 14},
		{key, E_Subkeys, 15},
	}
	for _, table := range tables {
		result := generateSubkeys(table.key, permutationChoice1[:], permutationChoice2[:], keyShiftRotations[:])
		if result[table.i] != table.expected_subkeys[table.i] {
			t.Errorf("generate subkeys, expected:%048b, got:%048b", table.expected_subkeys[0], result[0])
		}
	}
}

func TestInitialPermutation(t *testing.T) {
	tables := []struct {
		input    uint64
		expected uint64
	}{
		{
			0b0000000100100011010001010110011110001001101010111100110111101111,
			0b1100110000000000110011001111111111110000101010101111000010101010,
		},
	}
	for _, table := range tables {
		result := permuteBlockUniversal(table.input, initialPermutationTable[:], 64)
		if result != table.expected {
			t.Errorf("Initial Permutation. expected:%064b, got:%064b", table.expected, result)
		}
	}
}

func TestSplit(t *testing.T) {
	tables := []struct {
		input      uint64
		expected_L uint32
		expected_R uint32
	}{
		{
			0b1100110000000000110011001111111111110000101010101111000010101010,
			0b11001100000000001100110011111111,
			0b11110000101010101111000010101010,
		},
	}
	for _, table := range tables {
		result_L, result_R := splitNr(table.input, 64)
		if result_L != table.expected_L || result_R != table.expected_R {
			t.Errorf("TestSplit expected:(%064b, %064b), got: (%064b, %064b)", table.expected_L, table.expected_R, result_L, result_R)
		}
	}
}

func TestIteration(t *testing.T) {
	tables := []struct {
		input_L    uint32
		input_R    uint32
		expected_L uint32
		expected_R uint32
		subkey     uint64
	}{
		{
			input_L: 0b11001100000000001100110011111111,
			input_R: 0b11110000101010101111000010101010,

			subkey: 0b000110110000001011101111111111000111000001110010,

			expected_L: 0b11110000101010101111000010101010,
			expected_R: 0b11101111010010100110010101000100,
		},
	}
	for _, table := range tables {
		L1, R1 := des_Iteration(table.input_L, table.input_R, table.subkey, expansionTable[:], p_table_test[:])
		if L1 != table.expected_L {
			t.Errorf("\nLeft TestIteration :\nexpected:\t(%032b)\ngot:\t\t(%032b)", table.expected_L, L1)
		}
		if R1 != table.expected_R {
			t.Errorf("\nRight TestIteration:\nexpected:\t(%032b)\ngot:\t\t(%032b)", table.expected_R, R1)
		}
	}
}

func TestFFunction(t *testing.T) {
	tables := []struct {
		r        uint32
		k        uint64
		expected uint32
	}{
		{
			r:        0b11110000101010101111000010101010,
			k:        0b000110110000001011101111111111000111000001110010,
			expected: 0b00100011010010101010100110111011,
		},
	}
	for _, table := range tables {
		result := fFunction(table.r, table.k, expansionTable[:], p_table_test[:])
		if result != table.expected {
			t.Errorf("fFunction, expected:%032b, got:%032b", table.expected, result)
		}
	}
}

func TestExpansion(t *testing.T) {
	tables := []struct {
		input    uint64
		expected uint64
	}{
		{
			0b11110000101010101111000010101010,
			0b011110100001010101010101011110100001010101010101,
		},
	}
	for _, table := range tables {
		result := permuteBlockUniversal(table.input, expansionTable[:], 32)
		if result != table.expected {
			t.Errorf("Expansion Permutation. expected:%064b, got:%064b", table.expected, result)
		}
	}
}

func TestSboxPermutation(t *testing.T) {
	tables := []struct {
		input    uint64
		expected uint32
	}{
		{
			input:    0b011000010001011110111010100001100110010100100111,
			expected: 0b01011100100000101011010110010111,
		},
	}
	for _, table := range tables {
		result := sBoxPerm(table.input, sBox)
		if result != table.expected {
			t.Errorf("sBoxPerm, expected:%032b, got:%032b", table.expected, result)
		}
	}
}

func binAddition(a, b uint64) uint64 {
	return a ^ b
}

func TestBinAddition(t *testing.T) {
	tables := []struct {
		er       uint64
		k        uint64
		expected uint64
	}{
		{
			0b011110100001010101010101011110100001010101010101, //er
			0b000110110000001011101111111111000111000001110010, //k
			0b011000010001011110111010100001100110010100100111, //erk
		},
	}
	for _, table := range tables {
		result := binAddition(table.er, table.k)
		if result != table.expected {
			t.Errorf("binAddition, expected:%064b, got:%064b", table.expected, result)
		}
	}
}

func TestEncryption(t *testing.T) {
	tables := []struct {
		inputText uint64
		key       uint64
		expected  uint64
	}{
		{
			inputText: 0x0123456789ABCDEF,
			key:       0b0001001100110100010101110111100110011011101111001101111111110001,
			expected:  0x85E813540F0AB405,
		},
	}
	for _, table := range tables {
		result := desEncryption(table.inputText, table.key, true)
		if result != table.expected {
			t.Errorf("\ntestedEncryption\nexpected:\t%016x,\ngot:\t\t%016x", table.expected, result)
		}
	}
}

func TestDecrypt(t *testing.T) {
	tables := []struct {
		inputText uint64
		key       uint64
		expected  uint64
	}{
		{
			inputText: 0x85E813540F0AB405,
			key:       0b0001001100110100010101110111100110011011101111001101111111110001,
			expected:  0x0123456789ABCDEF,
		},
	}
	for _, table := range tables {
		result := desEncryption(table.inputText, table.key, false)
		if result != table.expected {
			t.Errorf("\ntestedEncryption\nexpected:\t%016x,\ngot:\t\t%016x", table.expected, result)
		}
	}
}

func TestStringSplit(t *testing.T) {
	tables := []struct {
		input    string
		expected []uint64
	}{
		{
			input:    "aaaabbbb",
			expected: []uint64{0x6161616162626262},
		},
		{
			input:    "aaaabb",
			expected: []uint64{0x6161616162623030},
		},
		{
			input:    "aaaabbbbaab",
			expected: []uint64{0x6161616162626262, 0x6161623030303030},
		},
		{
			input:    "HelloWorld",
			expected: []uint64{0x48656c6c6f576f72, 0x6c64303030303030},
		},
	}
	for _, table := range tables {
		result := splitStringBlocks(table.input)
		if result[len(result)-1] != table.expected[len(table.expected)-1] {
			t.Errorf("testedFunction(%s), expected:%0x, got:%0x", table.input, table.expected, result)
		}
	}
}

func TestStringMerge(t *testing.T) {
	tables := []struct {
		input          []uint64
		expectedString string
	}{
		{
			input:          []uint64{0x6161616162626262},
			expectedString: "aaaabbbb",
		},
		{
			input:          []uint64{0x6161616162623030},
			expectedString: "aaaabb00",
		},
		{
			input:          []uint64{0x6161616162626262, 0x6130303030303030},
			expectedString: "aaaabbbba0000000",
		},
	}
	for _, table := range tables {
		result := mergeStringBlocks(table.input)
		if result != table.expectedString {
			t.Errorf("testedFunction(%0x), expected:%s, got:%s", table.input[0], table.expectedString, result)
		}
	}
}
