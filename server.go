package main

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

func startServer() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	// Listen for an incoming connection.
	// Only accept one connection, the server itself.
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		os.Exit(1)
	}
	// Handle connections in a new goroutine.
	go readFromConn(conn)
}

// Handles incoming requests.
func readFromConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)
	for {
		line, err := tp.ReadLine()
		fmt.Println(line)
		if err != nil {
			fmt.Println("Closing connection")
			return
		}
	}
}
