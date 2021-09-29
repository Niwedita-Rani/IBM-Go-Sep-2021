package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileReader, err := os.Open("data1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer fileReader.Close()
	bufferedReader := bufio.NewReader(fileReader)
	counter := 0
	for {

		line, err := bufferedReader.ReadString('.')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		counter++
		fmt.Println(counter, line)

	}
}
