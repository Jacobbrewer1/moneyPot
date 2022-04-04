package controllers

import (
	"fmt"
	"github.com/Jacobbrewer1/moneypot/config"
	"github.com/Jacobbrewer1/moneypot/helper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/sheets/v4"
	"log"
	"os"
	"text/tabwriter"
	"time"
)

func SheetsSetup() *sheetService {
	// Create a JWT configurations object for the Google service account
	conf := &jwt.Config{
		Email:        *config.JsonConfigVar.RemoteConfig.Secrets.GoogleSheetCredentials.ClientEmail,
		PrivateKey:   []byte(*config.JsonConfigVar.RemoteConfig.Secrets.GoogleSheetCredentials.PrivateKey),
		PrivateKeyID: *config.JsonConfigVar.RemoteConfig.Secrets.GoogleSheetCredentials.PrivateKeyId,
		TokenURL:     *config.JsonConfigVar.RemoteConfig.Secrets.GoogleSheetCredentials.TokenUri,
		Scopes: []string{
			helper.SheetsScope,
		},
	}

	client := conf.Client(oauth2.NoContext)

	// Create a service object for Google sheets
	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
	return (*sheetService)(srv)
}

func (s *sheetService) GetSheetData() {
	// Pull the data from the sheet
	resp, err := s.Spreadsheets.Values.Get(*config.JsonConfigVar.RemoteConfig.SheetId, helper.SheetsRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	// Display pulled data
	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		log.Println("Spreadsheet is now:")
		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		for _, row := range resp.Values {
			var rowText string
			for _, column := range row {
				rowText = rowText + fmt.Sprintf("%v\t", column)
			}
			fmt.Fprintln(w, rowText)
		}
		w.Flush()
	}
}

func (s *sheetService) PostSheetData(l LoggingLine) {

	if l.Date.IsZero() {
		l.Date = time.Now().UTC()
	}
	if l.MoneyGoing == "" {
		l.MoneyGoing = "N/a"
	}
	if l.MoneyFrom == "" {
		l.MoneyFrom = "N/a"
		l.Amount = 0 - l.Amount
	}

	var d = sheets.ValueRange{
		Values: [][]interface{}{{l.Date.Format(time.RFC1123), l.Amount, l.MoneyFrom, l.MoneyGoing}},
	}

	_, err := s.Spreadsheets.Values.Append(*config.JsonConfigVar.RemoteConfig.SheetId, helper.SheetsRange, &d).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		log.Fatalln(err)
	}

	go s.GetSheetData()
}
