package Studienarbeit_src

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	Tcp_conn()
}
func Tcp_conn() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		if string(message) == "HALLO" {
			fmt.Print("tschüss")
		}
		fmt.Print("Message Received test:", string(message))
		// sample process for string received
		newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
		if message == "exit" {
			break
		}

	}
}
