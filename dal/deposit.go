package dal

import "log"

func DepositMoney(depo float64) {
	currentAmount, err := readAmount()
	if err != nil {
		log.Println(err)
	}
	log.Printf("current db amount = %v\n", currentAmount)
	if err := updateAmount(depo + currentAmount); err != nil {
		log.Println(err)
	}
	log.Printf("new value = %v\n", depo+currentAmount)
}
