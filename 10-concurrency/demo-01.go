package main

import (
	"fmt"
)

func main() {
	go print("Hello")
	print("World")
	fmt.Println("End of main")
	//time.Sleep(1 * time.Second)
	/*
		var input string
		fmt.Scanln(&input)
	*/
	//runtime.Gosched()
}

func print(str string) {
	fmt.Println(str)
}
