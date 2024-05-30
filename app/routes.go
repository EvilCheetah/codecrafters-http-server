package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"path/filepath"
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
	echo_text := strings.TrimPrefix(request.URL.Path, ECHO_PATH)

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

func handle_get_file(connection net.Conn, request *http.Request) {
	filename := strings.TrimPrefix(request.URL.Path, FILE_PATH)

	file_path := filepath.Clean(
		filepath.Join(*WEB_ROOT_PATH, filename),
	)

	file_stats, err := os.Stat(file_path)
	if errors.Is(err, os.ErrNotExist) {
		response := http.Response{
			ProtoMajor: 1,
			ProtoMinor: 1,
			StatusCode: http.StatusNotFound,
		}

		response.Write(connection)
		return
	} else if err != nil {
		fmt.Println(err.Error())

		handle_internal_server_error(connection, request)
		return
	}

	file, err := os.Open(file_path)
	if err != nil {
		fmt.Println(err.Error())

		handle_internal_server_error(connection, request)
		return
	}
	defer file.Close()

	response := http.Response{
		ProtoMajor:    1,
		ProtoMinor:    1,
		StatusCode:    http.StatusOK,
		ContentLength: file_stats.Size(),
		Header:        make(http.Header),
		Body:          io.NopCloser(file),
	}

	response.Header.Set("Content-Type", "application/octet-stream")

	response.Write(connection)
}

func handle_post_file(connection net.Conn, request *http.Request) {
	filename := strings.TrimPrefix(request.URL.Path, FILE_PATH)

	file, err := os.Create(
		filepath.Clean(filepath.Join(*WEB_ROOT_PATH, filename)),
	)
	if err != nil {
		fmt.Println("Unable to CREATE a file")

		handle_internal_server_error(connection, request)
		return
	}

	_, err = io.Copy(file, request.Body)
	if err != nil {
		fmt.Println("Unable to SAVE a file")

		handle_internal_server_error(connection, request)
		return
	}

	response := http.Response{
		ProtoMajor: 1,
		ProtoMinor: 1,
		StatusCode: http.StatusCreated,
	}

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

func handle_internal_server_error(connection net.Conn, request *http.Request) {
	response := http.Response{
		ProtoMajor: 1,
		ProtoMinor: 1,
		StatusCode: http.StatusInternalServerError,
	}

	response.Write(connection)
}
