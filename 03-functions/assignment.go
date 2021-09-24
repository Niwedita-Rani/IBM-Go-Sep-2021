package main

import "fmt"

func main() {
	var userChoice, n1, n2, result int
	for {
		userChoice = getUserChoice()
		switch userChoice {
		case 1, 2, 3, 4:
			fmt.Println("Enter the two numbers :")
			fmt.Scanf("%d %d", &n1, &n2)
			switch userChoice {
			case 1:
				result = add(n1, n2)
				fmt.Println("Addition of two numbers is :", result)
			case 2:
				result = subtract(n1, n2)
				fmt.Println("Subtraction of two numbers is :", result)
			case 3:
				result = multiply(n1, n2)
				fmt.Println("Multiplication of two numbers is :", result)
			case 4:
				result = divide(n1, n2)
				fmt.Println("Division of two numbers is :", result)
			}
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

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
