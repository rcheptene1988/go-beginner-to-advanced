package main

import "fmt"

type Rectangle interface {
	Perimeter() float64
	Area() float64
}

// Calculate struct
type Calculate struct {
	length float64
	width  float64
}

// Perimeter method
func (c Calculate) Perimeter() float64 {
	return (c.width + c.length) * 2
}

// Area method
func (c Calculate) Area() float64 {
	return c.width * c.length
}

func main() {

	var r Rectangle
	r = Calculate{20, 10}

	fmt.Println("Perimeter of rectangle: ", r.Perimeter())
	fmt.Println("Area of rectangle: ", r.Area())
}
