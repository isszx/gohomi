package main

import "fmt"

// Task #1
func factorial(i uint) uint {
	sum := 1
	for z := 1; z <= int(i); z++ {
		sum *= z
	}
	return uint(sum)
}

// Task #2
type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	return Point{s.start.x + int(s.a), s.start.y + int(s.a)}
}

func (s Square) Perimeter() uint {
	return s.a * 4
}

func (s Square) Area() uint {
	return s.a * s.a
}

func main() {
	fmt.Println("Factorial 10 =", factorial(5))

	s := Square{Point{1, 1}, 5}
	fmt.Println("Square end point:", s.End())
	fmt.Println("Square perimeter =", s.Perimeter())
	fmt.Println("Square area =", s.Area())
}
