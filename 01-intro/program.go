package main

import "fmt"

func main() {

	/*
		var msg string
		msg = "Hello, World!"
	*/

	/*
		var msg string = "Hello, World!"
	*/
	/*
		var msg = "Hello, World!" //type inference
	*/
	//the following syntax is applicable only in a function (Not at the package level)
	msg := "Hello, World!"
	fmt.Printf("%T = %q\n", msg, msg)

	//constants
	const myVar int = 100

	//multiple variable
	/*
		var x int
		var y int
		var result int
		var template string

		x = 100
		y = 200
		result = x + y
		template = "Adding %d and %d results in %d\n"
	*/

	/*
		var x, y, result int
		var template string
	*/

	/*
		var (
			x, y, result int
			template     string
		)

		x = 100
		y = 200
		result = x + y
	*/

	/*
		var (
			x, y, result int    = 100, 200, 0
			template     string = "Adding %d and %d results in %d\n"
		)
		result = x + y
	*/

	x, y, template := 100, 200, "Adding %d and %d results in %d\n"
	result := x + y
	fmt.Printf(template, x, y, result)

	//type conversion
	var i int = 100
	var f float64
	f = float64(i)
	fmt.Println(f)

}
