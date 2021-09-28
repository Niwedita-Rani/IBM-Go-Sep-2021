package main

import (
	"fmt"
	"time"
)

/*
the "hello" & "world" should be printed alternatively
*/
func main() {
	ch1, ch2 := make(chan string), make(chan string)
	go print("hello", ch1, ch2)
	go print("world", ch2, ch1)
	ch1 <- "start"
	var input string
	fmt.Scanln(&input)
}

func print(s string, in, out chan string) {
	for i := 0; i < 5; i++ {
		<-in
		time.Sleep(time.Second * 2)
		println(s)
		out <- "done"
	}
}
