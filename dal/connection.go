package dal

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	db *sql.DB
)

func DbSetup() {
	go loopCon(db)
	db = connect()
	runUpgrades()
}

func loopCon(d *sql.DB) {
	for {
		log.Println("connecting to db via go routine")
		d = connect()
		log.Println("waiting 2 minutes to connect again")
		time.Sleep(time.Minute * 2)
	}
}

func connect() *sql.DB {
	dbTemp, err := sql.Open("mysql", "root:Dummy1!SQL@tcp(127.0.0.1:1433)/moneypot?timeout=2s")
	if err != nil {
		log.Fatal(err)
	}

	dbTemp.SetConnMaxLifetime(time.Minute * 3)
	if err := dbTemp.Ping(); err != nil {
		log.Fatal(err)
	}
	return dbTemp
}

func runUpgrades() {
	if findFile("./dal/upgrades") {
		log.Println("upgrades detected")
		items, err := ioutil.ReadDir("./dal/upgrades/")
		if err != nil {
			log.Println(err)
		}

		for _, dir := range items {
			if dir.IsDir() {
				log.Printf("%v is dir\n", dir.Name())
				scripts, err := ioutil.ReadDir(fmt.Sprintf("./dal/upgrades/%v/", dir.Name()))
				if err != nil {
					log.Fatal(err)
				}

				for _, s := range scripts {
					file, err := ioutil.ReadFile(fmt.Sprintf("./dal/upgrades/%v/%v", dir.Name(), s.Name()))
					if err != nil {
						log.Println(err)
					}

					sqlCmd := string(file)

					if err := db.Ping(); err != nil {
						log.Fatal(err)
					}

					log.Printf("running upgrade /%v/%v", dir.Name(), s.Name())
					stmt, err := db.Prepare(sqlCmd)
					if err != nil {
						log.Fatal(err)
					}

					_, err = stmt.Exec()
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("upgrade /%v/%v completed successfully", dir.Name(), s.Name())
				}
			}
		}
	}
}

func findFile(path string) bool {
	abs, err := filepath.Abs(path)
	if err != nil {
		return false
	}
	log.Println(abs)

	file, err := os.Open(abs)
	if err != nil {
		return false
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	return true
}
