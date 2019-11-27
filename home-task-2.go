package main

import "math"

type Figure interface {
	area() float64
	perimeter() float64
}

type Square struct {
	width float64
}

func (s Square) area() float64 {
	return s.width * s.width
}

func (s Square) perimeter() float64 {
	return s.width * 4
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
