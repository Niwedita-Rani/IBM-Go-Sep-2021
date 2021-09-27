package main

import (
	"fmt"
	calc "modularity-demo/calculator" /* package alias */
	"modularity-demo/calculator/utils"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	color.Green(fmt.Sprintf("Add result = %d\n", calc.Add(10, 20)))
	color.Red(fmt.Sprintf("Subtract result = %d\n", calc.Subtract(10, 20)))

	log.Println(utils.IsEven(10))
	log.Println("Op Count = ", calc.GetOpCount())
}
