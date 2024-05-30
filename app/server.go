package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func HandleRequest(connection net.Conn) {
	defer connection.Close()

	request, err := http.ReadRequest(bufio.NewReader(connection))
	if err != nil {
		fmt.Println("Error reading the request", err.Error())
		return
	}

	switch path := request.URL.Path; {
	case path == "/":
		handle_root(connection, request)

	case strings.HasPrefix(request.URL.Path, "/echo/"):
		handle_echo(connection, request)

	default:
		handle_not_found(connection, request)
	}
}

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

	HandleRequest(connection)
}
