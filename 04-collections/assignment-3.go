package main

import "fmt"

var operations map[int]func(int, int) int = map[int]func(int, int) int{
	1: add,
	2: subtract,
	3: multiply,
	4: divide,
}

func main() {
	var userChoice, n1, n2, result int
	for {
		userChoice = getUserChoice()
		switch userChoice {
		case 1, 2, 3, 4:
			operFn := operations[userChoice]
			n1, n2 = getOperands()
			result = operFn(n1, n2)
			fmt.Println("Result is :", result)
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func getOperands() (int, int) {
	var n1, n2 int
	fmt.Println("Enter the two numbers :")
	fmt.Scanf("%d %d", &n1, &n2)
	return n1, n2
}

/* func getOper(userChoice int) func(int, int) int {
	switch userChoice {
	case 1:
		return add
	case 2:
		return subtract
	case 3:
		return multiply
	case 4:
		return divide
	}
	return func(x, y int) int { return 0 }
}
*/
func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func divide(x, y int) int {
	return x / y
}
func multiply(x, y int) int {
	return x * y
}

func getUserChoice() int {
	var userChoice int
	fmt.Println("Menu")
	fmt.Println("================================================================")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("5. Exit")
	fmt.Println("Enter the choice :")
	fmt.Scanf("%d", &userChoice)
	return userChoice
}
