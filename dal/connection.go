package dal

import (
	"database/sql"
	"fmt"
	"github.com/Jacobbrewer1/moneypot/config"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

// When calling this method, use the next line for defer db.close()
func connect() *sql.DB {
	db, err := sql.Open("mysql", *config.JsonConfigVar.RemoteConfig.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func Upgrade() {
	if findFile("./dal/upgrades") {
		log.Println("upgrades detected")
		items, err := ioutil.ReadDir("./dal/upgrades/")
		if err != nil {
			log.Println(err)
		}

		log.Println("connected to db")
		db := connect()
		defer db.Close()
		log.Println("db connected")

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
		log.Println("all upgrades completed")
		return
	}
	log.Println("no upgrades were detected")
	return
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
