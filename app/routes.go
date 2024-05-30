package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

func handle_root(connection net.Conn, request *http.Request) {
	response := http.Response{
		ProtoMajor: 1,
		ProtoMinor: 1,
		StatusCode: http.StatusOK,
	}

	response.Write(connection)
}

func handle_echo(connection net.Conn, request *http.Request) {
	echo_text := strings.TrimPrefix(request.URL.Path, "/echo/")

	fmt.Printf("Echoing: `%s`\n", echo_text)

	response := http.Response{
		ProtoMajor:    1,
		ProtoMinor:    1,
		StatusCode:    http.StatusOK,
		ContentLength: int64(len(echo_text)),
		Body:          io.NopCloser(bytes.NewBufferString(echo_text)),
		Header:        make(http.Header, 0),
	}

	response.Header.Set("Content-Type", "text/plain")

	response.Write(connection)
}

func handle_not_found(connection net.Conn, request *http.Request) {
	response := http.Response{
		ProtoMajor: 1,
		ProtoMinor: 1,
		StatusCode: http.StatusNotFound,
	}

	response.Write(connection)
}
