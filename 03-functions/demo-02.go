package main

import "fmt"

func main() {
	fmt.Println(sum(10))
	fmt.Println(sum(10, 20, 30))
	fmt.Println(sum(10, 20, 30, 40, 50))
	fmt.Println(sum(0))
}

func sum(x int, nos ...int) int {
	//nos => slice (array) of integers
	result := x
	for idx := 0; idx < len(nos); idx++ {
		result += nos[idx]
	}
	return result
}
