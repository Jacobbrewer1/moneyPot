package dal

import (
	"log"
	"sync"
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
