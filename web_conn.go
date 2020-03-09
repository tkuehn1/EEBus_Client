package Studienarbeit_src

import (
	// "fmt"
	// "io"
	"log"
	"net/http"
	"strings"
)

func web_conn(w http.ResponseWriter, req *http.Request) {
	ws, err := NewWebSocket(w, req)
	if err != nil {
		log.Println("Error creating websocket connection: %v", err)
		return
	}
	ws.On("message", func(e *Event) {
		log.Printf("Message reveived: %s", e.Data.(string))
		ws.Out <- (&Event{
			Name: "response",
			Data: strings.ToUpper(e.Data.(string)),
		}).Raw()
	})
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func Web_start() {
	http.Handle("/", http.FileServer(http.Dir("./assets")))
	http.HandleFunc("/ws", web_conn)

	err := http.ListenAndServeTLS(":7070", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
