package main

import (
	"html/template"
	"log"
	"net/http"
)

func handleFilePath() {
	log.Println("parsing templates")
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	templates = template.Must(template.New("").ParseGlob("./assets/templates/*.html"))
	log.Println("Files parsed successfully")
}

func depositMoneyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("deposit money request received")
	amount := r.FormValue("depositMoneyInput")
	log.Println(amount)
}

func home(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index", nil); err != nil {
		log.Fatal(err)
	}
}
