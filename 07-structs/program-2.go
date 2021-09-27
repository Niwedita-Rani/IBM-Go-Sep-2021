package main

import "fmt"

type ProductType struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

//version 1.0
/*
type Food struct {
	Product ProductType
	Expiry  string
}
*/

//version 2.0
type Food struct {
	ProductType
	Expiry string
	Name   string
}

func main() {
	//Version 1.0
	/*
		grapes := Food{Product: ProductType{Id: 101, Name: "Grapes", Cost: 60, Units: 40, Category: "Food"}, Expiry: "2 Days"}
		fmt.Println(grapes)
		fmt.Println(grapes.Product.Name)

		applyDiscount(&grapes.Product, 10)
		fmt.Println(grapes)
	*/

	//Version 2.0
	//grapes := Food{ProductType: ProductType{Id: 100, Name: "Grapes", Cost: 60, Units: 40, Category: "Food"}, Expiry: "2 Days"}
	grapes := &Food{
		ProductType{Id: 100, Name: "Grapes", Cost: 60, Units: 40, Category: "Food"},
		"2 Days",
		"Kabul Grapes",
	}
	fmt.Println(grapes)
	fmt.Println(grapes.ProductType.Name)
	fmt.Println(grapes.Name)

	//dummy := new(Food) //returns a pointer to the newly created object
	dummy := &Food{}
	fmt.Println(dummy)

	applyDiscount(&grapes.ProductType, 10)
	fmt.Println(grapes)

}

func applyDiscount(p *ProductType, discount float64) {
	p.Cost = p.Cost - (p.Cost * discount / 100)
}
