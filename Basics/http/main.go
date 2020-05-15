package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}

func main() {
	res, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("HTTP Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, res.Body)
}
