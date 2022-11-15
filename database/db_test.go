package database

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestDatabase(t *testing.T) {
	db := NewDB()
	tables := []struct {
		login    string
		password string
	}{
		{login: "John Cena", password: "MMA Champion"},
		{login: "Santa", password: "Merry Christmas"},
		{login: "Danu", password: "Nanu"},
		{login: "Alexandru", password: "Pretty bird"},
		{login: "Constantin", password: "Password123"},
	}
	for i, table := range tables {
		db.Add(table.login, table.password)
		if !db.CheckPassword(table.login, table.password) {
			hash1 := []byte(db.getPass(table.login))
			hash2 := sha256.Sum256([]byte(table.password + db.getSalt(table.login)))
			err1 := fmt.Sprintf("\nStored passwords\t%0x\nStored Salts\t%0x", db.passwordsHash, db.passwordsHash)
			err2 := fmt.Sprintf("\nTest %d\tLogin:%s\tPassword:%s\nSalt:\t%0x\nPassword hashes do not match\n\tPassHash:\t%0x\n\tStoredHash:\t%0x", i+1, table.login, table.password, db.getSalt(table.login), hash1, hash2)
			t.Errorf("%s%s", err1, err2)
		}
	}
}
