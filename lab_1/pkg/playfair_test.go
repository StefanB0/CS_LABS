package pkg

import "testing"

func TestPlayFairCypher(t *testing.T) {
	pfKey := "PLAYFIRBCDEGHKMNOQSTUVWXZ"
	tables := []struct {
		message, phrase string
		decrypt         bool
		expected        string
	}{
		{"Hello World!", pfKey, false, "KGYVRVVQGRCZ"},
		{"apple", pfKey, false, "YLLAKU"},
		{"calculus", pfKey, false, "BYYRVPXN"},
		{"KGYVRVVQGRCZ", pfKey, true, "HELXLOWORLDX"}, //decrypt
		{"YLLAKU", pfKey, true, "APPLEX"},
		{"BYYRVPXN", pfKey, true, "CALCULUS"},
	}
	for _, table := range tables {
		result := PlayFairCypher(table.message, table.phrase, table.decrypt)
		if result != table.expected {
			t.Errorf("Encrypted message %s is wrong, expected:%s, got:%s", table.message, table.expected, result)
		}
	}
}

func TestPlayFairKey(t *testing.T) {
	tables := []struct {
		input    string
		expected string
	}{
		{"Play Fair", "PLAYFIRBCDEGHKMNOQSTUVWXZ"},
		{"Gravity Falls", "GRAVITYFLSBCDEHKMNOPQUWXZ"},
		{"No Mans Sky", "NOMASKYBCDEFGHILPQRTUVWXZ"},
		{"Hello woman", "HELOWMANBCDFGIKPQRSTUVXYZ"},
		{"Hello World", "HELOWRDABCFGIKMNPQSTUVXYZ"},
	}

	for _, table := range tables {
		result := PlayFairKey(table.input)
		if result != table.expected {
			t.Errorf("PlayFair Key of phrase %s is wrong, expected:%s, got:%s", string(table.input), string(table.expected), string(result))
		}
	}
}

func TestFairSwitch(t *testing.T) {
	pfKey := "PLAYFIRBCDEGHKMNOQSTUVWXZ"
	tables := []struct {
		pair     string
		key      string
		decrypt  bool
		expected string
	}{
		{"PZ", pfKey, false, "FU"},
		{"RS", pfKey, false, "CO"},
		{"IM", pfKey, false, "DE"},
		{"RQ", pfKey, false, "BO"},
		{"AB", pfKey, false, "BH"},
		{"TZ", pfKey, false, "ZF"},
		{"OQ", pfKey, false, "QS"},
		{"CD", pfKey, false, "DI"},
		{"LF", pfKey, false, "AP"},
		{"RO", pfKey, false, "GV"},
		{"GV", pfKey, false, "OL"}, 
		{"FU", pfKey, true, "PZ"}, // Decrypt
		{"CO", pfKey, true, "RS"},
		{"DE", pfKey, true, "IM"},
		{"BO", pfKey, true, "RQ"},
		{"BH", pfKey, true, "AB"},
		{"ZF", pfKey, true, "TZ"},
		{"QS", pfKey, true, "OQ"},
		{"DI", pfKey, true, "CD"},
		{"AP", pfKey, true, "LF"},
		{"GV", pfKey, true, "RO"},
		{"OL", pfKey, true, "GV"},
	}
	for _, table := range tables {
		result := fairSwitch(table.pair, table.key, table.decrypt)
		if result != table.expected {
			t.Errorf("Switch pair: %s, expected: %s, got: %s", table.pair, table.expected, result)
		}
	}
}

func TestAdjustMessage(t *testing.T) {
	tables := []struct {
		input    string
		expected string
	}{
		{"Hello", "HELXLO"},
		{"apple dvd", "APPLEDVD"},
		{"@Home RUN!", "HOMERUNX"},
		{"Hello woman!", "HELXLOWOMANX"},
	}

	for _, table := range tables {
		result := adjustMessage(table.input)
		if result != table.expected {
			t.Errorf("Adjust message(%s), expected:%s, got:%s", string(table.input), string(table.expected), string(result))
		}
	}
}

func TestAddNonRepeatingString(t *testing.T) {
	tables := []struct {
		input    string
		expected string
	}{
		{"Play Fair", "PLAYFIR"},
		{"Gravity Falls", "GRAVITYFLS"},
		{"No Mans Sky", "NOMASKY"},
		{"Hello woman !", "HELOWMAN"},
	}

	for _, table := range tables {
		result := addNonRepeatingString("", table.input)
		if result != table.expected {
			t.Errorf("AddNonRepeatingString(%s), expected:%s, got:%s", string(table.input), string(table.expected), string(result))
		}
	}
}
