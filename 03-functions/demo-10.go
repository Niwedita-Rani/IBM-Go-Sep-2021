package main

import "fmt"

func main() {
	defer fmt.Println("Deferred from [@main] - 1")
	fmt.Println("Main started")
	f1()
	fmt.Println("Main completed")
}

func f1() {
	defer fmt.Println("Deferred from [@f1] - 1")
	defer fmt.Println("Deferred from [@f1] - 2")
	defer fmt.Println("Deferred from [@f1] - 3")
	f2()
	fmt.Println("f1 completed")
}

func f2() {
	defer fmt.Println("Deferred from [@f2] - 1")
	defer fmt.Println("Deferred from [@f2] - 2")
	defer fmt.Println("Deferred from [@f2] - 3")
	fmt.Println("f2 completed")
	panic("Something bad happend")

}

func processFile() {
	//open the file
	// defer - close the file

	//for {read until EOF}

	//get the line
	/*
		if err {
			return
		}
	*/

	//parse the line
	/*
		if err {
			return
		}
	*/

	//process the data
	/*
		if err {
			return
		}
	*/
	//}}

	//return result
}
