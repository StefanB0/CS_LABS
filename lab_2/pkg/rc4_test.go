package pkg

import "testing"

var lorem_ipsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

func TestRC4(t *testing.T) {
	k1:= "Hello World"
	k2:= "I believe I can fly"
	k3:= "I believe I can touch the sky"
	tables := []struct {
		plaintex []byte
		key      []byte
	}{
		{[]byte(lorem_ipsum), []byte(k1)},
		{[]byte(lorem_ipsum), []byte(k2)},
		{[]byte(lorem_ipsum), []byte(k3)},
	}
	for _, table := range tables {
		encryptedText := RC4Encrypt(table.plaintex, table.key)
		decryptedText := RC4Decrypt(encryptedText, table.key)
		if string(table.plaintex) != string(decryptedText) {
			t.Errorf("Decrypted text differs from plaintext")
		}
	}
}
