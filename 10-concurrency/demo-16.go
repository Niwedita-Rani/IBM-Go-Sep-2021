package main

import (
	"time"
)

func main() {
	ch := make(chan int)

	go fibonacci(ch)

	for no := range ch {
		println(no)
	}
}

func fibonacci(ch chan int) {
	/*
		stop := make(chan bool)
		go func() {
			time.Sleep(20 * time.Second)
			stop <- true
		}()
	*/
	stop := time.After(20 * time.Second)
	x, y := 0, 1
	for {
		select {
		case <-stop:
			close(ch)
			return
		case ch <- x:
			x, y = y, x+y
			time.Sleep(time.Millisecond * 500)
		}
	}
}
