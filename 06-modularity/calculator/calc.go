package calculator

import "fmt"

var opCount int

//package initializing
func init() {
	fmt.Println("calculator package initialized")
}

func GetOpCount() int {
	return opCount
}
