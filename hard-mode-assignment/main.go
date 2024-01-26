package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// os.Open() returns a file and an error
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Errorf("Error: ", err)
		os.Exit(1)
	}

	// the file struct has the Write function so it satisfies the Reader Interface
	io.Copy(os.Stdout, f)
}
