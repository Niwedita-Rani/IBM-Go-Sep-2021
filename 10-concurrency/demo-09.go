package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println(result)
	fmt.Println("End of main")
}

func add(x, y int, ch chan int) {
	result := x + y
	time.Sleep(time.Second * 5)
	ch <- result
}
