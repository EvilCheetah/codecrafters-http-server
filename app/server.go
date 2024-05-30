package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

var WEB_ROOT_PATH *string

const (
	ECHO_PATH       = "/echo/"
	FILE_PATH       = "/files/"
	USER_AGENT_PATH = "/user-agent"

	FILE_SIZE_LIMIT = 100 * MegaByte
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

	case strings.HasPrefix(request.URL.Path, ECHO_PATH):
		handle_echo(connection, request)

	case strings.HasPrefix(request.URL.Path, FILE_PATH) && request.Method == "GET":
		handle_get_file(connection, request)

	case strings.HasPrefix(request.URL.Path, FILE_PATH) && request.Method == "POST":
		handle_post_file(connection, request)

	case strings.HasPrefix(request.URL.Path, USER_AGENT_PATH):
		handle_user_agent(connection, request)

	default:
		handle_not_found(connection, request)
	}
}

func main() {
	WEB_ROOT_PATH = flag.String(
		"directory",
		"",
		"WebRoot Directory for `GET /files/<filename>`",
	)
	flag.Parse()

	listen, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		connection, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go HandleRequest(connection)
	}
}
