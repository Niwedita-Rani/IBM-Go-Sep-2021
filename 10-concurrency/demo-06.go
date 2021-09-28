package main

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

//communicate by sharing memory [NOT advicable in Golang]
type OpCounter struct {
	count int
	sync.Mutex
}

func (op *OpCounter) Increment() {
	op.Lock()
	{
		op.count++
	}
	op.Unlock()
}

var opCounter = &OpCounter{}

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

	fmt.Println("No of operations = ", opCounter.count)
	fmt.Println("End of main")
}

func add(x, y int, wg *sync.WaitGroup) {
	result := x + y
	fmt.Printf("Adding %d and %d = %d\n", x, y, result)
	opCounter.Increment()
	wg.Done()
}
