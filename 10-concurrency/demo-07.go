package main

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	wg.Add(1)
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println(result)
	wg.Wait()
	fmt.Println("End of main")
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result
	wg.Done()
}
