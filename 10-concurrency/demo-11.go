package main

import (
	"fmt"
	"time"
)

/*
the "hello" & "world" should be printed alternatively
*/
func main() {
	go print("hello")
	go print("world")
	var input string
	fmt.Scanln(&input)
}

func print(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 2)
		println(s)
	}
}
