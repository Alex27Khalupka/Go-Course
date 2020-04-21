package main

import "fmt"

// Represents a point with (x, y) coordinates
type Point struct {
	x, y int
}

// Represents a square with the coordinate
// of the upper left point(variable start) and side length
// equal to variable a
type Square struct {
	start Point
	a     uint
}

// Returns the the bottom left
// point of the square
func (s Square) End() Point {
	var end Point
	end.x = s.start.x + int(s.a)
	end.y = s.start.y + int(s.a)
	return end
}

// Returns the perimeter of the square
func (s Square) Perimeter() uint {
	return s.a * 4

}

// Returns the area of the square
func (s Square) Area() uint {
	return s.a * s.a
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
