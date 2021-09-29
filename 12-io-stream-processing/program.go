package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	fileWg := &sync.WaitGroup{}
	processWg := &sync.WaitGroup{}
	dataCh := make(chan int)
	oddCh := make(chan int)
	evenCh := make(chan int)
	oddSumCh := make(chan int)
	evenSumCh := make(chan int)
	fileWg.Add(2)
	go source("data1.dat", dataCh, fileWg)
	go source("data2.dat", dataCh, fileWg)

	processWg.Add(4)
	go splitter(dataCh, evenCh, oddCh, processWg)
	go sum(oddCh, oddSumCh, processWg)
	go sum(evenCh, evenSumCh, processWg)
	go merger(oddSumCh, evenSumCh, "result.dat", processWg)
	fileWg.Wait()
	close(dataCh)
	processWg.Wait()
	fmt.Println("Done")

}

func merger(oddSumCh, evenSumCh chan int, resultFile string, wg *sync.WaitGroup) {
	file, err := os.Create(resultFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		file.Close()
		close(oddSumCh)
		close(evenSumCh)
		wg.Done()
	}()

	for idx := 0; idx < 2; idx++ {
		select {
		case oddSum := <-oddSumCh:
			file.WriteString(fmt.Sprintf("Odd Total = %d\n", oddSum))
		case evenSum := <-evenSumCh:
			file.WriteString(fmt.Sprintf("Even Total = %d\n", evenSum))
		}
	}
}

func sum(noCh chan int, resultCh chan int, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	var sum int
	for no := range noCh {
		sum += no
	}
	resultCh <- sum
}

func splitter(dataCh chan int, evenCh, oddCh chan int, wg *sync.WaitGroup) {
	defer func() {
		close(evenCh)
		close(oddCh)
		wg.Done()
	}()
	for no := range dataCh {
		if no%2 == 0 {
			evenCh <- no
		} else {
			oddCh <- no
		}
	}
}

func source(fileName string, dataCh chan int, wg *sync.WaitGroup) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		file.Close()
		wg.Done()
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		no, err := strconv.Atoi(scanner.Text())
		if err == nil {
			dataCh <- no
		}
	}

}
