package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("division by zero error")

func main() {
	/*
		result, err := divide(5, 0)
		if err != nil {
			fmt.Println("Something went wrong!", err)
			return
		}
		fmt.Println("Result:", result)
	*/

	_, err := divide(5, 0)
	if err == DivideByZeroError {
		fmt.Println("Divide by zero error")
	}
	if err != nil {
		fmt.Println("Division failed")
		return
	}
	fmt.Println("Division succeeded!")
	fmt.Printf("Error = %v\n", err)
}

/* func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero error")
	}
	return x / y, nil
} */

func divide(x, y int) (result int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	result = x / y
	return
}
