package main

import (
	"github.com/Jacobbrewer1/moneypot/dal"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func init() {
	log.Println("initializing logging")
	//log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.Println("logging initialized")
}

func main() {
	dal.DbSetup()
	handleFilePath()

	r := mux.NewRouter()

	log.Println("listening...")

	r.HandleFunc("/", home)

	http.Handle("/", r)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
