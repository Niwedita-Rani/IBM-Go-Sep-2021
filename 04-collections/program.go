package main

import "fmt"

func main() {
	//Arrays
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	var nos [5]int = [...]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	//Iterating an array (using the indexer)
	for idx := 0; idx < len(nos); idx++ {
		fmt.Println(nos[idx])
	}

	//Iterating an array (using range)
	for _, val := range nos {
		fmt.Println(val)
	}

	//creating a copy of an array
	var nos2 [5]int = nos
	fmt.Println(nos2)
	nos2[0] = 200

	fmt.Println(nos, nos2)
}
