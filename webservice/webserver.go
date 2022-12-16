package webservices

import (
	"CS/Cryptography/database"
	"net/http"
)

const (
	ADMIN = "ADMIN"
	NORMAL_USER = "NORMAL_USER"
)

type WebServer struct {
	userDB *database.Database
	sessionDB UserTokens
	roleDB map[string]string
	http.Server
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (s *WebServer) StartServer() {
	s.userDB = database.NewDB()
	s.roleDB = make(map[string]string)
	s.CreateUsers()
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
