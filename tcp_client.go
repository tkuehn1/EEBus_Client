package Studienarbeit_src

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Tcp_client() {
	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text+"\n")
		// listen for reply
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}
		for {
			//check if the transmittet message is correctly received
			if message != text {
				fmt.Fprintf(conn, text+"\n")
			} else if message == text {
				if message == "exit" {
					fmt.Println("Close Connection\n")
					conn.Close()
					fmt.Print("Last Message from Server: " + message)
					return
				}
				fmt.Print("Last Message from Server: " + message)
				break
			}
		}
	}
}
