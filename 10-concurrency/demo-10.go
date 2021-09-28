package main

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

type OpCounter struct {
	count  int
	result int
}

func main() {

	ch := make(chan OpCounter)
	go add(100, 200, ch, 1)
	go add(1000, 200, ch, 2)
	go add(100, 2000, ch, 3)
	go add(1000, 2000, ch, 4)

	counter := 0
	results := []int{}
	for idx := 0; idx < 4; idx++ {
		fmt.Printf("[opId-%d] Attempting to read result\n", idx)
		opCounter := <-ch
		fmt.Printf("[opId-%d] completed reading result\n", idx)
		counter += opCounter.count
		results = append(results, opCounter.result)
	}
	fmt.Println("No of operations = ", counter)
	fmt.Println("Result = ", results)
	fmt.Println("End of main")
}

func add(x, y int, ch chan OpCounter, opId int) {
	result := x + y
	opCounter := OpCounter{count: 1, result: result}
	fmt.Printf("[opId-%d] Attempting to write result\n", opId)
	ch <- opCounter
	fmt.Printf("[opId-%d] Completed writing the result\n", opId)
}
