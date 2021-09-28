package main

import (
	"sync"
	"time"
)

/*
the "hello" & "world" should be printed alternatively
*/
func main() {
	wg := &sync.WaitGroup{}
	ch1, ch2 := make(chan string, 1), make(chan string, 1)
	wg.Add(2)
	go print("hello", ch1, ch2, wg)
	go print("world", ch2, ch1, wg)
	ch1 <- "start"
	wg.Wait()
}

func print(s string, in, out chan string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		<-in
		time.Sleep(time.Millisecond * 500)
		println(s)
		out <- "done"
	}
	wg.Done()
}
