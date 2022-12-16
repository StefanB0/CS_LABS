package main

import "CS/Cryptography/webservice"

func main() {
	web := webservice.NewWebServer()
	web.StartServer()
}
