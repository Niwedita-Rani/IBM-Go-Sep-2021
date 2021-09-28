package main

import (
	"fmt"
	"sync"
)

var result int
var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	wg.Add(1)
	go func() {
		result = add(100, 200)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int) int {
	return x + y

}
