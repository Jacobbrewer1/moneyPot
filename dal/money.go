package dal

import (
	"github.com/Jacobbrewer1/moneypot/controllers"
	"github.com/Jacobbrewer1/moneypot/helper"
	"log"
	"sync"
	"time"
)

func DepositMoney(depo float64, w ...*sync.WaitGroup) {
	if w != nil {
		defer w[0].Done()
	}

	currentAmount, err := ReadAmount()
	if err != nil {
		log.Println(err)
	}
	log.Printf("current db amount = %v\n", currentAmount)
	if err := updateAmount(depo + currentAmount); err != nil {
		log.Println(err)
	}
	log.Printf("new value = %v\n", depo+currentAmount)
}

func WithdrawMoney(amt float64, w ...*sync.WaitGroup) {
	if w != nil {
		defer w[0].Done()
	}

	currentAmount, err := ReadAmount()
	if err != nil {
		log.Println(err)
	}
	if amt > currentAmount {
		log.Println("not enough money to withdraw")
		return
	}
	if err := updateAmount(currentAmount - amt); err != nil {
		log.Println(err)
	}
	log.Printf("new value = %v\n", currentAmount-amt)
}

func SyncLoop() {
	for {
		SyncWithSheets()
		diff := helper.GetNext30MinTime().Sub(time.Now())
		log.Printf("waiting %v until syncing with sheets again\n", diff)
		time.Sleep(diff)
	}
}

func SyncWithSheets() {
	client := controllers.SheetsSetup()
	sheetTotal, err := client.GetTotal()
	if err != nil {
		log.Println(err)
		return
	}
	dbTotal, err := ReadAmount()
	if err != nil {
		log.Println(err)
		return
	}
	if sheetTotal == dbTotal {
		log.Println("db and sheets are in sync")
		return
	}
	log.Printf("sheets value (£%.2f) and db value (£%.2f) do not match\n", sheetTotal, dbTotal)
	log.Printf("updating to sheets value of £%.2f\n", sheetTotal)
	if err := updateAmount(sheetTotal); err != nil {
		log.Println(err)
		return
	}
	return
}
