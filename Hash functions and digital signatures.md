# Asymmetric Ciphers

### Course: Cryptography & Security
### Author: Boicu Stefan

----

## Objectives:

* Learn the theory and principles of hashing and digital signatures
* Use an appropriate hashing algorithms to store passwords in a local DB.
* Use an asymmetric cipher to implement a digital signature process for a user message.

### Key settings

First, install the golang language on your system. [Link for download](https://go.dev/learn/)

This laboratory work also has test coverage. Tests prove that the cipher work as intended. To verify that all the tests run correctly run.

`$go test ./signature/ -v`

# Hashing

## Theory

Hashing functions are a special kind of algorithm that take some data and output a fixed length number. They are defined by two characteristics, irreversibility and uniqueness. It is impossible to get the original data from a hash, and even a small change to the initial data will drastically change the hash. As such it is almost impossible to tamper with data and receive an identical hash afterwards. As such hashes are very useful in verifying the authencity and integrity of data. 

Given the a password it is trivial to find its hash, but you cannot do the reverse. As such most databses store only the hashes of passwords, and not the passwords themselves, so in case of a breach they woulnd't be leaked. Unfortunately criminals might try to crack the password using a dictionary attack. If they hashed a lot of common passwords, they could match them to the stolen password hashes and find them out this way. To prevent this an additional technique called salting is used. This involves adding a salt to the password before hashing and making a dictionary attack much harder.

## Implementation

Simple database structure. Stores **hashed** passwords and the salts used for hashing, as well as a dictionary that binds logins to passwords. The Database is append-only.

```go
type database struct {
	loginDictionary map[string]int
	passwordsHash   []string
	salts           []string
}
```

When you add a login-password pair to the database, it automatically generates a random salt and hashes the password, then it stores the hash and the salt in memory.

```go
func (db *database) Add(login, password string) {
	hpassword, salt := saltHashPassword(password)

	db.loginDictionary[login] = len(db.passwordsHash)
	db.passwordsHash = append(db.passwordsHash, string(hpassword[:]))
	db.salts = append(db.salts, string(salt))
}
```

The salt is a randomly generated array of 32 bytes, which is the same length as the final hash. Which makes even short passwords more secure against decoding the hash.

```go
func saltHashPassword(s string) (hash [32]byte, salt []byte) {
	salt = make([]byte, SALT_SIZE)
	for i := 0; i < len(salt); i++ {
		salt[i] = byte(rand.Intn(256))
	}
	hash = sha256.Sum256(append([]byte(s), salt...))
	return hash, salt
}
```

The last function of the database is to check if a login-password pair is valid

```go

func (db *database) CheckPassword(login, pass string) bool {
	pass1 := db.getPass(login)
	pass2 := sha256.Sum256([]byte(pass + db.getSalt(login)))
	return pass1 == string(pass2[:])
}
```

# Digital signature

## Theory

Digital signatures are used to verify the validity of online communication. The message is hashed and then encrypted using an asymetric algorithm. The result is the signature which is appended to the message and sent over the network. The receiver uses the public key to decrypt the hash, then compares it to the hash of the message. If the hashes are equal and the public key is valid, the receiver can be assured that the message was not tampered on.

## Implementation

I sign a message by first hashing the message, then encrypt the hash with RSA

```go
func signMessage(plaintext string, publicKey rsa.Key) (signature []uint64) {
	hash := sha256.Sum256([]byte(plaintext))
	stringHash := string(hash[:])
	binaryHash := rsa.SplitString(stringHash)
	signature = rsa.RSAEncryption(binaryHash, publicKey)
	return
}
```

I check a signature by decrypting the hash, hashing the rest of the message, and then check if the hashes are equal.

```go
func verifySignature(message string, signature []uint64, privateKey rsa.Key) bool {
	signatureHash := rsa.MergeToString(rsa.RSAEncryption(signature, privateKey))
	messageHash := sha256.Sum256([]byte(message))
	return signatureHash == string(messageHash[:])
}
```

## Conclusions

As part of this laboratory work I learned about hashing, digital signatures and how to implement them.