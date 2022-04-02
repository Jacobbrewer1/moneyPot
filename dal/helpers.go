package dal

import "log"

const (
	CurrencyDbIdValue = 0
	updateSql = "UPDATE `moneypot`.`Money` SET `amount` = ? WHERE (`id` = ?);"
)

func updateAmount(amt float64) error {
	sql, err := db.Prepare(updateSql)
	if err != nil {
		return err
	}
	defer sql.Close()

	query, err := sql.Query(amt, CurrencyDbIdValue)
	if err != nil {
		return err
	}
	log.Println(query)
	return nil
}
