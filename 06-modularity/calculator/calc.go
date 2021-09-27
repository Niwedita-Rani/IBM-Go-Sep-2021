package calculator

import "fmt"

//private
var opCount int

//package initializing
func init() {
	fmt.Println("calculator package initialized")
}

func GetOpCount() int {
	return opCount
}
