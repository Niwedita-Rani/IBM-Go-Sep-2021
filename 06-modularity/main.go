package main

import (
	"fmt"
	calc "modularity-demo/calculator" /* package alias */
	"modularity-demo/calculator/utils"
)

func main() {
	fmt.Println(calc.Add(10, 20))
	fmt.Println(calc.Subtract(10, 20))
	fmt.Println(utils.IsEven(10))
	fmt.Println("Op Count = ", calc.GetOpCount())
}
