package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	stop := make(chan bool)
	go fibonacci(ch, stop)

	go func() {
		fmt.Println("Hit ENTER to stop!")
		var input string
		fmt.Scanln(&input)
		stop <- true
	}()

	for no := range ch {
		println(no)
	}
}

func fibonacci(ch chan int, stop chan bool) {
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
