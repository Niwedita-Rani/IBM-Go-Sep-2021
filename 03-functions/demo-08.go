package main

import "fmt"

func main() {
	/*
		go f1()
		f2()
		f3()
	*/

	for idx := 0; idx < 10; idx++ {
		go func(x int) {
			printNo(x)
		}(idx)
	}
	fmt.Println("hit ENTER to exit")
	var input string
	fmt.Scanln(&input)
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}

func f3() {
	fmt.Println("f3 invoked")
}

func printNo(no int) {
	fmt.Println("Number : ", no)
}
