# Topic: Web Authentication & Authorisation.

### Course: Cryptography & Security
### Author: Vasile Drumea

----

## Overview

&ensp;&ensp;&ensp; Authentication & authorization are 2 of the main security goals of IT systems and should not be used interchangibly. Simply put, 
during authentication the system verifies the identity of a user or service, and during authorization the system checks the access rights, 
optionally based on a given user role.

&ensp;&ensp;&ensp; There are multiple types of authentication based on the implementation mechanism or the data provided by the user. 
Some usual ones would be the following:
- Based on credentials (Username/Password);
- Multi-Factor Authentication (2FA, MFA);
- Based on digital certificates;
- Based on biometrics;
- Based on tokens.

&ensp;&ensp;&ensp; Regarding authorization, the most popular mechanisms are the following:
- Role Based Access Control (RBAC): Base on the role of a user;
- Attribute Based Access Control (ABAC): Based on a characteristic/attribute of a user.


## Objectives:
1. Take what you have at the moment from previous laboratory works and put it in a web service / serveral web services.
2. Your services should have implemented basic authentication and MFA (the authentication factors of your choice).
3. Your web app needs to simulate user authorization and the way you authorise user is also a choice that needs to be done by you.
4. As services that your application could provide, the classical ciphers could be used . Basically the user would like to get access and use the classical ciphers, but they need to authenticate and be authorized. 

## Implementation description:

The webservice and all related files are in the `webservice` folder. The service provided is logging into the system and the the encryption and decryption of classical ciphers. To be more specific

- Caesar cipher
- Play Fair cipher
- Viginere cipher

The main taks of this laboratory work was the implementation of authentification and authorization functionality. This was achieved by allowing the user to log in with a login and password. At that point the user needs to check their email where they received a number code. They then need to send a request with the code to receive an authorization token. This token lets them use classical ciphers according to their **ROLE**. A normal user can only access the caesar cipher while an admin user can access all three ciphers.

When using the cipher API they must send a token along with the request. The webserver checks the token validity, its expiration time and the user role associated with it. If all checks are successful it continues and fulfills the user's request. Alternatively it sends back and error message.

## Code implementation

I used the following code to send email messages for 2FA, the email and password for the mail server are in an `.env` file which is **NOT** uploaded to github

```golang
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
```

When the user first sends the login password pair, only they are checked and an email is sent with the code. If the login or password is wrong an error message is returned

```go
if !s.userDB.CheckPassword(loginReq.Login, loginReq.Password) {
	w.Write([]byte("Invalid name or password"))
	return
}

code := rand.Intn(1000-100) + 100
s.codeDB[loginReq.Login] = code
codeString := strconv.Itoa(code)
s.sendEmail(codeString, "stfnbcx@gmail.com")
w.Write([]byte("A code was sent to your email"))

```

Afterwards the login, password and code is checked for validity, and an authorization token is returned

```go
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
```

Each http api endpoint checks for authorization before performing an operation. Here is the code that handles the caesar cipher. As can be seen, both admin and normal users can access it.

```go
if !s.sessionDB.Verify(caesarRequest.Token) {
	w.Write([]byte("user not authorized"))
	return
}

if s.checkRole(caesarRequest.Token.Login, []string{NORMAL_USER, ADMIN}) {
	w.Write([]byte("user has insufficient privileges"))
}
```

## Conclusions:
In this laboratory work I learned and implemented my own webserver that handles both authentification and authorization. I learned about the distinction between authentification and authorization as well as the different ways of authorization, including RBAC and ABAC and methods of authorization.

I implemented an email-based form of 2FA and integrated it with the rest of the system.
