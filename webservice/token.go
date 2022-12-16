package webservice

import (
	"log"
	"math/rand"
	"time"
)

const (
	EXPIRE_DURATION = time.Minute * 30
)

type UserToken struct {
	Token  string    `json:"token"`
	Login  string    `json:"login"`
	Expire time.Time `json:"expire"`
}

type UserTokens []UserToken

func (ut *UserTokens) Exists(token string) bool {
	for _, userToken := range *ut {
		if userToken.Token == token {
			return true
		}
	}
	return false
}

func (ut *UserTokens) isExpired(token string) bool {
	for _, userToken := range *ut {
		if time.Now().After(userToken.Expire) {
			return true
		}
	}
	return false
}

func (ut *UserTokens) Verify(uToken UserToken) bool {
	if ut.isExpired(uToken.Login) {
		log.Println("token expired")
		return false
	}

	if !ut.Exists(uToken.Token) {
		log.Println("token not found")
		return false
	}

	for _, userToken := range *ut {
		if userToken.Login == uToken.Login && userToken.Token == uToken.Token {
			return true
		}
	}
	log.Println("token not found")
	return false
}

func (ut *UserTokens) Add(uToken UserToken) {
	*ut = append(*ut, uToken)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	rand.Seed(time.Now().Unix())

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateToken(_login string) *UserToken {
	return &UserToken{
		Login:  _login,
		Token:  randSeq(10),
		Expire: time.Now().Truncate(time.Second).Add(EXPIRE_DURATION),
	}
}
