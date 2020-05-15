package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println("Area of shape is:", s.getArea())
}

func main() {
	s := square{sideLength: 5}
	printArea(s)

	t := triangle{height: 3, base: 3}
	printArea(t)
}
