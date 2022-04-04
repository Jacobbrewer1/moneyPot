package controllers

import (
	"google.golang.org/api/sheets/v4"
	"time"
)

type (
	sheetService sheets.Service

	LoggingLine struct {
		Date       time.Time
		Amount     float64
		MoneyFrom  string
		MoneyGoing string
	}
)
