package main

import "fmt"

func main() {
	//create a array of 10 numbers
	//write a function that will take the numbers array as argument and returns a new array
	// the new array should have the data populated as follows:
	//original = [3,1,4,2,5]
	//new array = ["odd", "odd", "even", "even", "odd"]

	var nos [10]int = [...]int{7, 1, 4, 2, 3, 8, 5, 6, 9, 10}
	oddEvenList := transform(nos)
	fmt.Println(nos, oddEvenList)
}

func transform(nos [10]int) [10]string {
	var oddEvenList [10]string
	for idx, value := range nos {
		if value%2 == 0 {
			oddEvenList[idx] = "even"
		} else {
			oddEvenList[idx] = "odd"
		}
	}
	return oddEvenList
}
