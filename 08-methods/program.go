package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

type Food struct {
	Product
	Expiry string
	Name   string
}

func main() {
	var pen = Product{1, "Pen", 1.99, 100, "Stationery"}
	//fmt.Println(Format(pen))
	fmt.Println(pen.Format())

	//applyDiscount(&pen, 10)
	//(&pen).applyDiscount(10)
	pen.applyDiscount(10)
	fmt.Println(pen.Format())

	var food = Food{Product{2, "Bread", 2.99, 100, "Food"}, "2020-01-01", "Bread"}
	//applyDiscount(&food.Product, 10)
	food.applyDiscount(10)
	fmt.Println(food.Format())
}

//as a function
/*
func Format(product Product) string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s\n", product.Id, product.Name, product.Cost, product.Units, product.Category)
}
*/

//as a method
func (product Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s\n", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

//as a function
/*
func applyDiscount(p *Product, discount float64) {
	p.Cost = p.Cost - (p.Cost * discount / 100)
}
*/

//as a method
func (p *Product) applyDiscount(discount float64) {
	p.Cost = p.Cost - (p.Cost * discount / 100)
}
