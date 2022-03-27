package main

import (
	"html/template"
	"log"
	"net/http"
)

func handleFilePath() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	templates = template.Must(template.New("").ParseGlob("./assets/templates/*.html"))
	log.Println("Files parsed successfully")
}

func home(w http.ResponseWriter, r *http.Request) {

}
