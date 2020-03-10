package Studienarbeit_src

import (
	"encoding/json"
	// "fmt"
	// "io"
	"log"
	"net/http"
)

type ServerResponse struct {
	Status string `json:"status"`
}

func web_conn(w http.ResponseWriter, req *http.Request) {
	ws, err := NewWebSocket(w, req)
	if err != nil {
		log.Println("Error creating websocket connection: %v", err)
		return
	}
	raw, _ := json.Marshal(ServerResponse{Status: "Success"})

	ws.On("message", func(e *Event) {
		req, _ := json.Marshal(e.Data)
		log.Printf("Message reveived: %s", req)
		ws.Out <- raw
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
