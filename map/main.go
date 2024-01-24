package main

import (
	"fmt"
)

func main() {
	// Maps store data
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"blue":  "#0000ff",
	}
	// deleting the red hax from the colors map
	// delete(colors, "red")

	printMap(colors)
}

// iterating through a map
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex for", color, "is", hex)
	}
}
