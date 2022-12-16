package webservice

import (
	"CS/Cryptography/database"
	"math/rand"
	"net/http"
	"time"
)

const (
	ADMIN       = "ADMIN"
	NORMAL_USER = "NORMAL_USER"
)

type WebServer struct {
	userDB    *database.Database
	sessionDB UserTokens

	roleDB map[string]string
	codeDB map[string]int

	http.Server
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (s *WebServer) StartServer() {
	rand.Seed(time.Now().UnixMicro())
	
	s.userDB = database.NewDB()
	s.roleDB = make(map[string]string)
	s.codeDB = make(map[string]int)
	s.CreateUsers()
	s.readCredentials()
	s.inithandlers()
}

func (s *WebServer) CreateUsers() {
	s.userDB.Add("Stefan", "supersecret")
	s.roleDB["Stefan"] = ADMIN

	s.userDB.Add("Bernard", "12345")
	s.roleDB["Bernard"] = NORMAL_USER
}

func (s *WebServer) checkRole(user string, target []string) bool {
	for _, role := range target {
		if user == role {
			return true
		}
	}
	return false
}
