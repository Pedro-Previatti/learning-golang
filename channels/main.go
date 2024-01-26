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

	// channels can only communicate one type
	c := make(chan string) // creating a channel of strings

	for _, link := range links {
		go checkTrafic(link, c)
	}

	// infinitly checking status
	for {
		go checkTrafic(<-c, c) // blocking call
	}
}

// checking requests are slow, because it waits a response for the previous request
// checking trafic from link
func checkTrafic(s string, c chan string) {
	_, err := http.Get(s) // blocking call
	fmt.Println("Checking status for:", s)

	// if the get request's status is not 200 err is returned not null so site is probably down
	if err != nil {
		fmt.Println("	Site status: DOWN")
		c <- s
		return
	}

	fmt.Println("Site status: UP")
	c <- s
}
