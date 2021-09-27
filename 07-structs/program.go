package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category CategoryType
}

type CategoryType struct {
	Name        string
	Description string
}

func main() {
	//var p Product
	p := Product{Id: 100, Name: "Pen", Cost: 5, Units: 100, Category: CategoryType{Name: "Stationary", Description: "Writing Instruments"}}
	//p := Product{100, "Pen", 10, 100, "Stationary"}
	fmt.Println(p)

	pPtr := &p
	//fmt.Println((*pPtr).Name)
	//Dereferencing a pointer is NOT required to access the attributes
	fmt.Println(pPtr.Name)
	applyDiscount(&p, 10)
	fmt.Println(p)

	p1 := Product{Id: 100, Name: "Pen", Category: CategoryType{Name: "Stationary", Description: "Writing Instruments"}}
	p2 := Product{Id: 100, Name: "Pen", Category: CategoryType{Name: "Stationary", Description: "Writing Instruments"}}

	//comparison by value
	fmt.Println(p1 == p2)

	//comparison by reference
	fmt.Println(&p1 == &p2)
}

func applyDiscount(p *Product, discount float64) {
	p.Cost = p.Cost - (p.Cost * discount / 100)
}
