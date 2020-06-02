package EEBus_Client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	Server_PORT = "7070"
)

func StartTCP() {
	go Tcp_socket()
}

func Tcp_socket() {
	// Listen for incoming connections.
	l, err := net.Listen("tcp", "localhost"+":"+Server_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + "localhost" + ":" + Server_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		msg = strings.TrimRight(msg, "\r\n")
		if msg == "high" {
			Led("high")
		} else if msg == "low" {
			Led("low")
		}

		println(msg)

		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}

func Client(src string) {
	//connect to this socket
	conn, _ := net.Dial("tcp", src+":"+Server_PORT)
	for {
		//read in input from stdin

		text := Taster()
		//sed to socket
		fmt.Fprintf(conn, text+"\n")
		//listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server:" + message)
	}
}
