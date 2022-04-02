package main

import (
	"github.com/Jacobbrewer1/moneypot/dal"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
)

var templates *template.Template

func init() {
	log.Println("initializing logging")
	//log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.Println("logging initialized")
}

func getIpAddress() (string, string) {
	name, err := os.Hostname()
	if err != nil {
		log.Println(err)
		return "", ""
	}

	addrs, err := net.LookupIP(name)
	if err != nil {
		log.Println(err)
		return "", ""
	}

	var i4, i6 string
	for _, a := range addrs {
		if ipv6 := a.To16(); ipv6 != nil {
			i6 = ipv6.String()
		}
		if ipv4 := a.To4(); ipv4 != nil {
			i4 = ipv4.String()
		}
	}
	return i4, i6
}

func main() {
	dal.DbSetup()
	handleFilePath()

	r := mux.NewRouter()

	log.Println("listening...")

	r.HandleFunc("/", home)
	r.HandleFunc("/depositMoney", depositMoneyHandler).Methods(http.MethodPost)
	r.HandleFunc("/withdrawMoney", withdrawMoneyHandler).Methods(http.MethodPost)
	r.HandleFunc("/live/updates/amount", liveUpdates).Methods(http.MethodGet)

	http.Handle("/", r)
	ip4, ip6 := getIpAddress()
	log.Printf("listening at:\nIPV4: %v:8080\nIPV6: %v:8080", ip4, ip6)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
