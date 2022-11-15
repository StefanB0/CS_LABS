package classical

import "testing"

func TestCaesarPermEncrypt(t *testing.T) {
	tables := []struct {
		s        string
		key      int
		alph     []rune
		expected string
	}{
		{"Hello World", 3, PermutateAlphabet(3), "Edpps Qsgpa"},
		{"I will be back", 5, PermutateAlphabet(5), "Q vqtt wh wazm"},
		{"To be? Or not to be, that is the question.", -2, PermutateAlphabet(5), "Cz pa? Zv kzc cz pa, cqtc jh cqa dyahcjzk."},
	}
	for _, table := range tables {
		result := CaesarPermEncrypt(table.s, table.key, table.alph)
		if result != table.expected {
			t.Errorf("CaesarPermEncrypt(%s, Key:%d), expected:%s, got:%s", table.s, table.key, table.expected, result)
		}
	}
}

func TestCaesarPermDecrypt(t *testing.T) {
	tables := []struct {
		s        string
		key      int
		alph     []rune
		expected string
	}{
		{"Edpps Qsgpa", 3, PermutateAlphabet(3), "Hello World"},
		{"Q vqtt wh wazm", 5, PermutateAlphabet(5), "I will be back"},
		{"Cz pa? Zv kzc cz pa, cqtc jh cqa dyahcjzk.", -2, PermutateAlphabet(5), "To be? Or not to be, that is the question."},
	}
	for _, table := range tables {
		result := CaesarPermDecrypt(table.s, table.key, table.alph)
		if result != table.expected {
			t.Errorf("CaesarPermDecrypt(%s, Key:%d), expected:%s, got:%s", table.s, table.key, table.expected, result)
		}
	}
}
