package main

import "fmt"

func main() {
	// Maps store data
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"blue":  "#0000ff",
	}

	numbers := make(map[int]int)

	numbers[1], numbers[2] = 0, 1
	// numbers[2] = "1"
	// this will not run because the map is of int: int

	fmt.Println(colors)
	// fmt.Println(colors.red)

	fmt.Println(numbers)

}
