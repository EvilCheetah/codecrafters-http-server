package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	
	connection, err := listen.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	connection.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
}
