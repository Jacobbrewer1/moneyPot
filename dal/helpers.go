package dal

const (
	CurrencyDbIdValue = 0
	updateSql         = "UPDATE `moneypot`.`Money` SET `amount` = ? WHERE (`id` = ?);"
	readSql           = "SELECT `amount` FROM `moneypot`.`Money` WHERE (`id` = ?);"
)

func updateAmount(amt float64) error {
	sql, err := db.Prepare(updateSql)
	if err != nil {
		return err
	}
	defer sql.Close()

	_, err = sql.Exec(amt, CurrencyDbIdValue)
	if err != nil {
		return err
	}
	return nil
}

func ReadAmount() (float64, error) {
	sql, err := db.Prepare(readSql)
	if err != nil {
		return 0, err
	}

	rows, err := sql.Query(CurrencyDbIdValue)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var amount float64
	for rows.Next() {
		if err := rows.Scan(&amount); err != nil {
			return 0, err
		}
	}
	return amount, nil
}
