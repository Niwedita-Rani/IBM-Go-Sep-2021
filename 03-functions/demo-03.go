package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous function invoked")
	}()

	func(x, y int) {
		fmt.Println(x + y)
	}(100, 200)

	/*
		var fn = func() {
			fmt.Println("fn is invoked")
		}
		fn()
	*/
	var fn func()
	fn = func() {
		fmt.Println("fn is invoked")
	}
	fn()

	var calc func(int, int) int
	calc = func(x, y int) int {
		return x + y
	}
	fmt.Println(calc(100, 200))

	calc = func(x, y int) int {
		return x - y
	}
	fmt.Println(calc(100, 200))

}
