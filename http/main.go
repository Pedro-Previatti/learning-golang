package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// type logWriter will get a function
type logWriter struct{}

func main() {
	// http.Get() returns a pointer to a Response and an Error (the return status of the http get request)
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// making a one liner to get the same body response
	io.Copy(lw, resp.Body)
}

// writing a Write function so that logWriter struct gets the Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	// normal implementation of Write function
	fmt.Println(string(bs))

	// custom line to show the ammount of bytes writen
	fmt.Println("Number of bytes", len(bs))

	return len(bs), nil
}
