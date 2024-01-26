package main

func main() {
	s := square{
		sideLength: 62.5,
	}

	t := triangle{
		base:   69.8,
		height: 22.2,
	}

	printArea(s)
	printArea(t)
}
