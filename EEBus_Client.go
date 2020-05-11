package main

import (
	"EEBus_Client"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/ws", EEBus_Client.WebConn)
	err := http.ListenAndServeTLS(":7070", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	for true {
		time.Sleep(time.Second)
	}
}
