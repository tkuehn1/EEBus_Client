package main

import (
	"EEBus_Client"
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	for true {

		fmt.Println("Run Menue")
		fmt.Println("--------------------------")
		fmt.Println("Type 1 to use the websocket function")
		fmt.Println("--------------------------")
		fmt.Println("Type 2 to connect two EEBUS-Clients with each other")
		fmt.Println("--------------------------")
		fmt.Print("->")
		reader := bufio.NewReader(os.Stdin)
		text, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println(err)
			return
		}

		switch text {
		case '1':
			web()
		case '2':
			fmt.Println("Is this a taster ore LED? taster = 1 | LED = 2")
			reader1 := bufio.NewReader(os.Stdin)
			text, _, err := reader1.ReadRune()

			if err != nil {
				fmt.Println(err)
				return
			}

			switch text {
			case '1':
				taster()
			case '2':
				led()
			}

		default:
			fmt.Println("!!!!!No Valid Input detected please correct your Input!!!!!")
		}
	}

}

func web() {
	http.HandleFunc("/ws", EEBus_Client.WebConn)
	err := http.ListenAndServeTLS(":7070", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	for true {
		time.Sleep(time.Second)
	}
}

func taster() {
	go EEBus_Client.Multicast("taster")
	go EEBus_Client.Tcp_socket()
	for {
		if EEBus_Client.GetSrcString() != "" {
			break
		}
		time.Sleep(1 * time.Second)
	}
	EEBus_Client.Client(EEBus_Client.GetSrcString())
}

func led() {
	go EEBus_Client.Multicast("led")
	go EEBus_Client.Tcp_socket()
}
