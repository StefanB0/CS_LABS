package webservice

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

var (
	emailLogin string
	emailPass  string
)

func (s *WebServer) readCredentials() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	emailLogin = os.Getenv("emailLogin")
	emailPass = os.Getenv("emailPass")
}

func (s *WebServer) sendEmail(code, receiver string) {
	to := []string{receiver}
	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(code)

	// Authentication.
	auth := smtp.PlainAuth("", emailLogin, emailPass, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, emailLogin, to, message)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Email Sent Successfully!")
}
