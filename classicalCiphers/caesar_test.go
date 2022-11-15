package classical

import "testing"

func TestCaesarEncrypt(t *testing.T) {
	tables := []struct {
		input    string
		key      int
		expected string
	}{
		{"Hello World", 3, "Khoor Zruog"},
		{"I will be back", 5, "N bnqq gj gfhp"},
		{"To be? Or not to be, that is the question.", -2, "Rm zc? Mp lmr rm zc, rfyr gq rfc oscqrgml."},
	}
	for _, table := range tables {
		result := CaesarEncrypt(table.input, table.key)
		if result != table.expected {
			t.Errorf("CaesarEncrypt(%s, %d), expected:%s, got:%s", table.input, table.key, table.expected, result)
		}
	}
}

func TestCaesarDecrypt(t *testing.T) {
	tables := []struct {
		input    string
		key      int
		expected string
	}{
		{"Khoor Zruog", 3, "Hello World"},
		{"N bnqq gj gfhp", 5, "I will be back"},
		{"Rm zc? Mp lmr rm zc, rfyr gq rfc oscqrgml.", -2, "To be? Or not to be, that is the question."},
	}
	for _, table := range tables {
		result := CaesarDecrypt(table.input, table.key)
		if result != table.expected {
			t.Errorf("CaesarDecrypt(%s, %d), expected:%s, got:%s", table.input, table.key, table.expected, result)
		}
	}
}
