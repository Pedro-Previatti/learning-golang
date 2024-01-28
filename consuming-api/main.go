package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	url := "https://api.ipify.org"

	res, _ := http.Get(url)
	ip, _ := io.ReadAll(res.Body)

	os.Stdout.Write(ip)
}
