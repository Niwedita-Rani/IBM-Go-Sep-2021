package main

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

//communicate by sharing memory [NOT advicable in Golang]
var opCount int = 0
var mutex = &sync.Mutex{}

func main() {
	wg.Add(1)
	go add(100, 200, wg)

	wg.Add(1)
	go add(1000, 200, wg)

	wg.Add(1)
	go add(100, 2000, wg)

	wg.Add(1)
	go add(1000, 2000, wg)

	wg.Wait()

	fmt.Println("No of operations = ", opCount)
	fmt.Println("End of main")
}

func add(x, y int, wg *sync.WaitGroup) {
	result := x + y
	fmt.Printf("Adding %d and %d = %d\n", x, y, result)
	mutex.Lock()
	{
		opCount++
	}
	mutex.Unlock()
	wg.Done()
}
