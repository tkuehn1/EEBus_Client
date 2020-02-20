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
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		message = strings.TrimRight(message, "\r\n")
		// output message received

		fmt.Print("Message Received:", string(message))
		// send received string back to client
		conn.Write([]byte(message + "\n"))
		if message == "exit" {
			fmt.Println("Close Connection")
			conn.Close()
			ln.Close()
			return
		}

	}
}
