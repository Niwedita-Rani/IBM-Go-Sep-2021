/*
	write a goroutine 'fibonacci' that generates 15 fibonacci numbers and write them into the given channel.

	consume the data from the channel and print the numbers as and when they are produced (in main)
*/

package main

func main() {
	ch := make(chan int)
	go fibonacci(ch)
	for no := range ch {
		println(no)
	}
}

func fibonacci(ch chan int) {
	x, y := 0, 1
	for idx := 0; idx < 15; idx++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
