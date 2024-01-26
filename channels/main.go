package main

import (
	"fmt"
	"net/http"
)

func main() {
	// links to check if site is up or down
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	for _, link := range links {
		checkTrafic(link)
	}
}

// checking trafic from link
func checkTrafic(s string) {
	_, err := http.Get(s)
	// if the get request's status is not 200 err is returned not null so site is probably down
	if err != nil {
		fmt.Println("Status check for:", s, "DOWN")
		return
	}

	fmt.Println("Status check for:", s, "OK")
}
