package main

import (
	"bytes"
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

	response := http.Response{
		ProtoMajor:    1,
		ProtoMinor:    1,
		StatusCode:    http.StatusOK,
		ContentLength: int64(len(echo_text)),
		Body:          io.NopCloser(bytes.NewBufferString(echo_text)),
		Header:        make(http.Header),
	}

	response.Header.Set("Content-Type", "text/plain")

	response.Write(connection)
}

func handle_user_agent(connection net.Conn, request *http.Request) {
	user_agent := request.Header.Get("User-Agent")

	response := http.Response{
		ProtoMajor:    1,
		ProtoMinor:    1,
		StatusCode:    http.StatusOK,
		ContentLength: int64(len(user_agent)),
		Body:          io.NopCloser(bytes.NewBufferString(user_agent)),
		Header:        make(http.Header),
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
