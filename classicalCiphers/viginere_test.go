package classical

import "testing"

func TestViginereEncrypt(t *testing.T) {
	tables := []struct {
		message    string
		key      string
		expected string
	}{
		{"hello world", "Key", "rijvs uyvjn"},
		{"The little red ridding hood went to her grandma", "wolf", "Pvp qeheqa fpi nwoiebr mkco babe yk vpw cflszal"},
		{"To the moon!", "sun", "Li gzy zgia!"},
	}
	for _, table := range tables {
		result := ViginereEncrypt(table.message, table.key)
		if result != table.expected {
			t.Errorf("ViginereEncrypt(%s, Key:%s), expected:%s, got:%s", table.message, table.key, table.expected, result)
		}
	}
}

func TestViginereDecrypt(t *testing.T) {
	tables := []struct {
		message    string
		key      string
		expected string
	}{
		{"rijvs uyvjn", "Key", "hello world"},
		{"Pvp qeheqa fpi nwoiebr mkco babe yk vpw cflszal", "wolf", "The little red ridding hood went to her grandma"},
		{"Li gzy zgia!", "sun", "To the moon!"},
	}
	for _, table := range tables {
		result := ViginereDecrypt(table.message, table.key)
		if result != table.expected {
			t.Errorf("ViginereDecrypt(%s, Key:%s), expected:%s, got:%s", table.message, table.key, table.expected, result)
		}
	}
}
