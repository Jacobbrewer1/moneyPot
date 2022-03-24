package main

import (
	"github.com/Jacobbrewer1/moneypot/dal"
	"log"
	"net/http"
)

func init() {
	log.Println("initializing logging")
	//log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.Println("logging initialized")
}

func main() {
	dal.DbSetup()

	log.Println("listening...")



	http.ListenAndServe("localhost:8080", nil)
}
