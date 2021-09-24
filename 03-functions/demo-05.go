package main

import "fmt"

func main() {
	fn := getFn()
	str := fn()
	fmt.Println(str)

}

func getFn() func() string {
	return func() string {
		fmt.Println("fn is invoked")
		return "some dummy string"
	}
}
