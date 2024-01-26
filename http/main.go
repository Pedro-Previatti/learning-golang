package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// http.Get() returns a pointer to a Response and an Error (the return status of the http get request)
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// creating a byte slice to implement Read function
	bs := make([]byte, 99_999)

	// reading the body of Response from the http and storing it in the byte slice
	resp.Body.Read(bs)

	// now we are printing the html from the body of the response
	fmt.Println(string(bs))
}
