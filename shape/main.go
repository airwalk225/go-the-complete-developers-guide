package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}
type square struct {
	sideLength float64
}

func main() {
	t := triangle{base: 1.2,
		height: 3.0,
	}

	s := square{sideLength: 3.0}

	fmt.Println("The area of the triangle is:", printArea(t))
	fmt.Println("The are of the square is:", printArea(s))

}

func printArea(s shape) float64 {
	return s.getArea()
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
