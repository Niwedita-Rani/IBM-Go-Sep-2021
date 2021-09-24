package main

import "fmt"

func main() {
	/*
		fmt.Println("invocation started")
		println(add(42, 13))
		fmt.Println("invocation completed")
		fmt.Println("invocation started")
		println(subtract(42, 13))
		fmt.Println("invocation completed")
	*/
	wrappedAdd := wrap(add)
	wrappedSubtract := wrap(subtract)
	wrappedAdd(42, 13)
	wrappedSubtract(42, 13)
}

func add(x, y int) int {
	return x + y
}
func subtract(x, y int) int {
	return x - y
}

func wrap(oper func(int, int) int) func(x, y int) {
	return func(x, y int) {
		fmt.Println("invocation started")
		fmt.Println(oper(x, y))
		fmt.Println("invocation completed")
	}
}
