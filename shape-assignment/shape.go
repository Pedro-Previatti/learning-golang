package main

import "fmt"

type ishape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return .5 * t.base * t.height
}

func printArea(i ishape) {
	fmt.Println("Area of shape: ", i.getArea())
}
