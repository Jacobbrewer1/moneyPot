package main

import (
	"fmt"
	"github.com/Jacobbrewer1/moneypot/config"
	"github.com/Jacobbrewer1/moneypot/dal"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
)

var templates *template.Template

var ws = websocket.Upgrader{
	HandshakeTimeout:  0,
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	WriteBufferPool:   nil,
	Subprotocols:      nil,
	Error:             nil,
	CheckOrigin:       nil,
	EnableCompression: false,
}
var websocketConnection *websocket.Conn

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

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	sock, err := ws.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	websocketConnection = sock

	log.Println("websocket created")

	reader(sock)
}

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		log.Printf("client message:\ntype: %v\nlog: %v\n", messageType, string(p))

		amt, err := dal.ReadAmount()
		if err != nil {
			log.Fatalln(err)
		}

		if err := conn.WriteMessage(messageType, []byte(fmt.Sprintf("£%.2f", amt))); err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	if err := config.ReadConfig(); err != nil {
		log.Fatalln(err)
	}

	dal.DbSetup()
	go dal.SyncLoop()

	handleFilePath()

	r := mux.NewRouter()

	log.Println("listening...")

	r.HandleFunc("/", home)
	r.HandleFunc("/ws", wsEndpoint) // Web socket endpoint
	r.HandleFunc("/depositMoney", depositMoneyHandler).Methods(http.MethodPost)
	r.HandleFunc("/withdrawMoney", withdrawMoneyHandler).Methods(http.MethodPost)

	http.Handle("/", r)

	ip4, ip6 := getIpAddress()
	log.Printf("listening at:\nIPV4: %v:8443\nIPV6: %v:8443", ip4, ip6)
	if err := http.ListenAndServe(":8443", nil); err != nil {
		log.Fatal(err)
	}
}
