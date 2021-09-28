package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg *sync.WaitGroup = &sync.WaitGroup{}
	wg.Add(1)
	go print("Hello", wg)
	wg.Add(1)
	go print("World", wg)

	fmt.Println("reaching end of main")
	wg.Wait()
	fmt.Println("end of main")
}

func print(str string, wg *sync.WaitGroup) {
	fmt.Println(str)
	wg.Done()
}
