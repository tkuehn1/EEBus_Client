package main

import (
	"EEBus_Client"
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Run Menue")
		fmt.Println("--------------------------")
		fmt.Println("Type web to use the websocket function")
		fmt.Println("--------------------------")
		fmt.Println("Type taster to connect two EEBUS-Clients with each other")
		fmt.Println("--------------------------")
		fmt.Print("->")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		switch text {
		case "web":
			web()
		case "taster":
			fmt.Println("Is this a taster ore LED? taster = 1 | LED = 2")
			text2, _ := reader.ReadString('\n')
			text2 = strings.Replace(text2, "\n", "", -1)
			switch text2 {
			case "1":
				taster()
			case "2":
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
	go EEBus_Client.StartTCP()
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
	go EEBus_Client.StartTCP()
}
