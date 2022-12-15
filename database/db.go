package database

import (
	"crypto/sha256"
	"math/rand"
)

const SALT_SIZE = 32

type Database struct {
	loginDictionary map[string]int
	passwordsHash   []string
	salts           []string
}

func NewDB() *Database {
	db := new(Database)
	db.loginDictionary = make(map[string]int)
	db.passwordsHash = []string{}
	db.salts = []string{}
	return db
}

func (db *Database) getPass(login string) string {
	return db.passwordsHash[db.loginDictionary[login]]
}

func (db *Database) getSalt(login string) string {
	return db.salts[db.loginDictionary[login]]
}

func (db *Database) Add(login, password string) {
	hpassword, salt := saltHashPassword(password)

	db.loginDictionary[login] = len(db.passwordsHash)
	db.passwordsHash = append(db.passwordsHash, string(hpassword[:]))
	db.salts = append(db.salts, string(salt))
}

func (db *Database) CheckPassword(login, pass string) bool {
	pass1 := db.getPass(login)
	pass2 := sha256.Sum256([]byte(pass + db.getSalt(login)))
	return pass1 == string(pass2[:])
}

func saltHashPassword(s string) (hash [32]byte, salt []byte) {
	salt = make([]byte, SALT_SIZE)
	for i := 0; i < len(salt); i++ {
		salt[i] = byte(rand.Intn(256))
	}
	hash = sha256.Sum256(append([]byte(s), salt...))
	return hash, salt
}
