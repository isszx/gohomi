package main

import (
	"math"
	"testing"
)

func Test_FigureOne(t *testing.T) {
	var s Figure = Square{10}
	var c Figure = Circle{10}

	if s.area() != 10 * 10 {
		t.Error("Square area incorrect")
	}
	if s.perimeter() != 10 * 4 {
		t.Error("Square perimeter incorrect")
	}

	if c.area() != math.Pi * 10 * 10 {
		t.Error("Circle area incorrect")
	}

	if c.perimeter() != 2 * math.Pi * 10 {
		t.Error("Circle perimeter incorrect")
	}
}

func Test_FigureTwo(t *testing.T) {
	var s Figure = Square{}
	var c Figure = Circle{}

	if s.area() != 0 {
		t.Error("Square area incorrect")
	}
	if s.perimeter() != 0 {
		t.Error("Square perimeter incorrect")
	}

	if c.area() != 0 {
		t.Error("Circle area incorrect")
	}

	if c.perimeter() != 0 {
		t.Error("Circle perimeter incorrect")
	}
}