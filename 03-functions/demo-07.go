package main

import "fmt"

var counter int = 0

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	reset()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}

func increment() int {
	counter++
	return counter
}

func reset() {
	counter = 0
}
