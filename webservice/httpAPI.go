package webservice

import (
	classical "CS/Cryptography/classicalCiphers"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

type loginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Code     int    `json:"code"`
}

type caesarCipherRequest struct {
	Text    string `json:"text"`
	Key     int    `json:"key"`
	Decrypt string `json:"decrypt"`

	Token UserToken `json:"userToken"`
}

type viginereCipherRequest struct {
	Text    string `json:"text"`
	Key     string `json:"key"`
	Decrypt string `json:"decrypt"`

	Token UserToken `json:"userToken"`
}

type playfairCipherRequest struct {
	Text    string `json:"text"`
	Key     string `json:"key"`
	Decrypt string `json:"decrypt"`

	Token UserToken `json:"userToken"`
}

func (s *WebServer) inithandlers() {
	http.HandleFunc("/", getHello)
	http.HandleFunc("/login", s.loginHandler)
	http.HandleFunc("/caesar", s.caesarHandler)
	http.HandleFunc("/playfair", s.playfairHandler)
	http.HandleFunc("/viginere", s.viginereHandler)

	s.initListen()
}

func (s *WebServer) initListen() {
	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("Server closed \n")
	} else if err != nil {
		log.Printf("error starting server %s\n", err)
		os.Exit(1)
	}
}

func getHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func (s *WebServer) loginHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Server could not read request body: %s\n", err)
	}

	var loginReq loginRequest
	json.Unmarshal(reqBody, &loginReq)

	if !s.userDB.CheckPassword(loginReq.Login, loginReq.Password) {
		w.Write([]byte("Invalid name or password"))
		return
	}

	if loginReq.Code != 0 && loginReq.Code == s.codeDB[loginReq.Login] {
		token := generateToken(loginReq.Login)
		s.sessionDB.Add(*token)

		tokenJson, err := json.Marshal(*token)
		if err != nil {
			log.Printf("Server could not write token to JSON: %s\n", err)
		}
		w.Write(tokenJson)
		s.codeDB[loginReq.Login] = 0
		return
	}

	code := rand.Intn(1000-100) + 100
	s.codeDB[loginReq.Login] = code
	codeString := strconv.Itoa(code)
	s.sendEmail(codeString, "stfnbcx@gmail.com")
	w.Write([]byte("A code was sent to your email"))

}

func (s *WebServer) caesarHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Server could not read request body: %s\n", err)
	}

	var caesarRequest caesarCipherRequest
	json.Unmarshal(reqBody, &caesarRequest)

	if !s.sessionDB.Verify(caesarRequest.Token) {
		w.Write([]byte("user not authorized"))
		return
	}

	if s.checkRole(caesarRequest.Token.Login, []string{NORMAL_USER, ADMIN}) {
		w.Write([]byte("user has insufficient privileges"))
	}

	if caesarRequest.Decrypt == "false" {
		w.Write([]byte(classical.CaesarEncrypt(caesarRequest.Text, caesarRequest.Key)))
		return
	} else if caesarRequest.Decrypt == "true" {
		w.Write([]byte(classical.CaesarDecrypt(caesarRequest.Text, caesarRequest.Key)))
		return
	}
	w.Write([]byte("invalid request"))
	return
}

func (s *WebServer) playfairHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Server could not read request body: %s\n", err)
	}

	var playfairRequest playfairCipherRequest
	json.Unmarshal(reqBody, &playfairRequest)

	if !s.sessionDB.Verify(playfairRequest.Token) {
		w.Write([]byte("user not authorized"))
		return
	}

	if s.checkRole(playfairRequest.Token.Login, []string{ADMIN}) {
		w.Write([]byte("user has insufficient privileges"))
	}

	if playfairRequest.Decrypt == "false" {
		w.Write([]byte(classical.PlayFairCypher(playfairRequest.Text, playfairRequest.Key, false)))
		return
	} else if playfairRequest.Decrypt == "true" {
		w.Write([]byte(classical.PlayFairCypher(playfairRequest.Text, playfairRequest.Key, true)))
		return
	}
	w.Write([]byte("invalid request"))
	return
}

func (s *WebServer) viginereHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Server could not read request body: %s\n", err)
	}

	var viginereRequest viginereCipherRequest
	json.Unmarshal(reqBody, &viginereRequest)

	if !s.sessionDB.Verify(viginereRequest.Token) {
		w.Write([]byte("user not authorized"))
		return
	}

	if s.checkRole(viginereRequest.Token.Login, []string{ADMIN}) {
		w.Write([]byte("user has insufficient privileges"))
	}

	if viginereRequest.Decrypt == "false" {
		w.Write([]byte(classical.ViginereEncrypt(viginereRequest.Text, viginereRequest.Key)))
		return
	} else if viginereRequest.Decrypt == "true" {
		w.Write([]byte(classical.ViginereDecrypt(viginereRequest.Text, viginereRequest.Key)))
		return
	}
	w.Write([]byte("invalid request"))
	return
}
