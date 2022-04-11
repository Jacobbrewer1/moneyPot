package main

import (
	"fmt"
	"github.com/Jacobbrewer1/moneypot/config"
	"github.com/Jacobbrewer1/moneypot/controllers"
	"github.com/Jacobbrewer1/moneypot/dal"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
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
	log.Printf("deposit amount of %.2f received\n", amount)

	from := r.FormValue("moneyFrom")
	log.Printf("Money is from %v\n", from)

	var waiter sync.WaitGroup
	if config.JsonConfigVar.UseDB() {
		waiter.Add(1)
		go dal.DepositMoney(amount, &waiter)
	}

	go createLog(controllers.LoggingLine{
		Date:          time.Time{},
		Amount:        amount,
		MoneyInReason: from,
	})

	waiter.Wait()
	go pushAmount(websocketConnection)
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
	log.Printf("withdraw amount of %.2f received\n", amount)
	reason := r.FormValue("withdrawReason")
	log.Printf("withdrawal reason is %v\n", reason)

	var waiter sync.WaitGroup
	if config.JsonConfigVar.UseDB() {
		waiter.Add(1)
		go dal.WithdrawMoney(amount, &waiter)
	}

	go createLog(controllers.LoggingLine{
		Date:           time.Time{},
		Amount:         amount,
		MoneyOutReason: reason,
	})

	waiter.Wait()
	go pushAmount(websocketConnection)
}

func createLog(line controllers.LoggingLine) {
	client := controllers.SheetsSetup()
	client.PostSheetData(line)
}

func pushAmount(conn *websocket.Conn) {
	if conn == nil {
		log.Println("websocket connection dead")
		return
	}

	amount, err := dal.ReadAmount()
	if err != nil {
		log.Println(err)
	}
	log.Printf("updating live amount with %v\n", amount)
	if err := conn.WriteMessage(1, []byte(fmt.Sprintf("£%.2f", amount))); err != nil {
		if websocket.IsCloseError(err) {
			log.Println("websocket connection dead")
		} else {
			log.Fatalln(err)
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("home page request received")
	if err := templates.ExecuteTemplate(w, "index", nil); err != nil {
		log.Fatal(err)
	}
}
