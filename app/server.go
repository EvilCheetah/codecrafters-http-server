package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"
)

const (
	HTTP_STATUS_200_OK        = "HTTP/1.1 200 OK\r\n\r\n"
	HTTP_STATUS_404_NOT_FOUND = "HTTP/1.1 404 Not Found\r\n\r\n"
)

func HandleRequest(connection net.Conn) {
	defer connection.Close()

	request, err := http.ReadRequest(bufio.NewReader(connection))
	if err != nil {
		fmt.Println("Error reading the request", err.Error())
		return
	}

	if request.URL.Path == "/" {
		connection.Write([]byte(HTTP_STATUS_200_OK))
		return
	}

	connection.Write([]byte(HTTP_STATUS_404_NOT_FOUND))
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
