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
	wrapInvoke(add, 42, 13)
	wrapInvoke(subtract, 42, 13)
}

func add(x, y int) int {
	return x + y
}
func subtract(x, y int) int {
	return x - y
}

func wrapInvoke(oper func(int, int) int, x, y int) {
	fmt.Println("invocation started")
	println(oper(x, y))
	fmt.Println("invocation completed")
}
