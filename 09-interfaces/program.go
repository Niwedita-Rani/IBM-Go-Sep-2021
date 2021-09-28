package main

import (
	"fmt"
	"math"
)

//Sprint-1
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Sprint-2
type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

//Sprint-3
type ShapeWithArea interface {
	Area() float64
}

func PrintArea(shape ShapeWithArea) {
	fmt.Println("Area = ", shape.Area())
}

//Sprint-4
func PrintPerimeter( /*  */ ) {
	fmt.Println("Perimeter = ", c.Perimeter())
}

/*
implement whatever is needed
*/

func main() {
	c := Circle{Radius: 5}
	//fmt.Println("Area = ", c.Area())
	PrintArea(c)
	PrintPerimeter(c)

	r := Rectangle{Height: 5, Width: 10}
	//fmt.Println("Area = ", r.Area())
	PrintArea(r)
	PrintPerimeter(r)
}
