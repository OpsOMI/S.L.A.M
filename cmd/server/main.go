package main

import (
	"fmt"
	"log"
	"net"
	"time"

	runserver "github.com/OpsOMI/S.L.A.M/internal/app/server"
	"github.com/OpsOMI/S.L.A.M/internal/configs/server"
)

func main() {
	configs := server.LoadConfig(
		"./configs/server.yaml",
		"./env/real/.env.management",
		"./deployment/dev/.env",
		"./deployment/prod/.env",
	)
	runserver.Run(*configs)
}

func ExampleMain() {
	ln, err := net.Listen("tcp", "localhost:6666")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		// Accept an incoming connection
		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		// Handle Connection
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		conn.SetReadDeadline(time.Now().Add(time.Second))

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			} else {
				log.Println("Connection closed or error: ", err)
				break
			}
		}

		if n == 0 {
			log.Println("Connection closed by client")
			break
		}

		msg := string(buf[:n])
		fmt.Println("Received: ", msg)

		sendMsg := "Hello I'm Server"
		fmt.Println("Send: ", sendMsg)

		_, err = fmt.Fprintln(conn, sendMsg)
		if err != nil {
			log.Println("Error sending message: ", err)
			break
		}
	}
}
