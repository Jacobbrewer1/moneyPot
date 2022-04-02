package main

import (
	"fmt"
	"github.com/Jacobbrewer1/moneypot/dal"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func handleFilePath() {
	log.Println("parsing templates")
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	templates = template.Must(template.New("").ParseGlob("./assets/templates/*.html"))
	log.Println("Files parsed successfully")
}

func depositMoneyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("deposit money request received")
	formValue := r.FormValue("depositMoneyInput")
	if formValue == "" {
		log.Println("deposit value is nil")
		http.Error(w, "deposit value is nil", 0)
		return
	}
	amount, err := strconv.ParseFloat(formValue, 32)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 1)
		return
	}
	if amount < 0 {
		log.Println("amount is less than zero")
		log.Println("invalid amount")
		http.Error(w, "invalid amount received", 2)
	}
	log.Printf("deposit amount of %v received\n", amount)
	go dal.DepositMoney(amount)
}

func withdrawMoneyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("withdraw money request received")
	formValue := r.FormValue("withdrawMoneyInput")
	if formValue == "" {
		log.Println("withdraw value is nil")
		http.Error(w, "withdraw value is nil", 0)
		return
	}
	amount, err := strconv.ParseFloat(formValue, 32)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), 1)
		return
	}
	if amount < 0 {
		log.Println("amount is less than zero")
		log.Println("invalid amount")
		http.Error(w, "invalid amount received", 2)
	}
	log.Printf("withdraw amount of %v received\n", amount)
	go dal.WithdrawMoney(amount)
}

func liveUpdates(w http.ResponseWriter, r *http.Request) {
	amount, err := dal.ReadAmount()
	if err != nil {
		log.Println(err)
	}
	log.Printf("updating live amount with %v\n", amount)
	w.Write([]byte(fmt.Sprintf("£%.2f", amount)))
}

func home(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index", nil); err != nil {
		log.Fatal(err)
	}
}
