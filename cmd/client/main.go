package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Create a TCP Socket
	conn, err := net.Dial("tcp", "localhost:6666")
	if err != nil {
		log.Fatal(err)
	}

	// Send a message
	fmt.Println("Sending: Hello, I'm Client")
	fmt.Fprint(conn, "Hello, I'm Client")

	// read message
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(buf[:n]))
	conn.Close()
}
