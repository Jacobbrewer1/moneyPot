package dal

import "log"

func DepositMoney(depo float64)  {
	if err := updateAmount(depo); err != nil {
		log.Println(err)
	}
}
