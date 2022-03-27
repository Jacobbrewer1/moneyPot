package main

import (
	"github.com/Jacobbrewer1/moneypot/dal"
	"github.com/gorilla/mux"
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

	r := mux.NewRouter()

	log.Println("listening...")

	r.HandleFunc("/", home)

	http.ListenAndServe("localhost:8080", nil)
}
