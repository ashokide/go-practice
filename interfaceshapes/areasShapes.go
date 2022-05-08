package main

import "fmt"

type Shapes interface {
	area() float32
}

type square struct {
	a float32
}

type rectangle struct {
	l float32
	b float32
}

func (s square) area() {
	fmt.Println("Area of Square: ", s.a*s.a)
}

func (r rectangle) area() {
	fmt.Println("Area of Rectangle: ", r.l*r.b)
}

func main() {
	s := square{10.50}
	s.area()
	r := rectangle{12, 6}
	r.area()
}
