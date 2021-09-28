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

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

//Sprint-2
type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

//Sprint-3
type ShapeWithArea interface {
	Area() float64
}

func PrintArea(shape ShapeWithArea) {
	fmt.Println("Area = ", shape.Area())
}

//Sprint-4

type ShapeWithPerimeter interface {
	Perimeter() float64
}

func PrintPerimeter(shape ShapeWithPerimeter) {
	fmt.Println("Perimeter = ", shape.Perimeter())
}

//Sprint-5
func PrintDimensions(sd ShapeWithDimensions) {
	PrintArea(sd)
	PrintPerimeter(sd)
}

type ShapeWithDimensions interface {
	/* Area() float64
	Perimeter() float64 */
	ShapeWithArea
	ShapeWithPerimeter
}

//Sprint-5
type AreaCalculator interface {
	Area() float64
}

func calculateArea(ac AreaCalculator) string {
	return fmt.Sprintf("Calculated Area = %f", ac.Area())
}

func main() {
	c := Circle{Radius: 5}
	//fmt.Println("Area = ", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintDimensions(c)

	r := Rectangle{Height: 5, Width: 10}
	//fmt.Println("Area = ", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintDimensions(r)

	fmt.Println(calculateArea(c))
}
