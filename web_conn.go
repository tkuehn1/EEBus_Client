package Studienarbeit_src

import (
	// "fmt"
	// "io"
	"log"
	"net/http"
)

func web_conn(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func Web_start() {
	http.HandleFunc("/hello", web_conn)
	http.Handle("/", http.RedirectHandler("/hello", 302))
	err := http.ListenAndServeTLS(":7070", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
