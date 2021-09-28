package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func (p Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type Products []Product

func (products Products) String() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprintf("%v\n", product)
	}
	return result
}

func (products Products) Len() int {
	return len(products)
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

//for id
func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

//for name
type ByName struct {
	Products
}

func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

//for units
type ByUnits struct {
	Products
}

func (byUnits ByUnits) Less(i, j int) bool {
	return byUnits.Products[i].Units < byUnits.Products[j].Units
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	fmt.Println("Initial List")
	fmt.Println(products)

	fmt.Println("Sorting [Default (by id)]")
	sort.Sort(products)
	fmt.Println(products)

	fmt.Println("Sorting [by name]")
	sort.Sort(ByName{products})
	fmt.Println(products)

	fmt.Println("Sorting [by units]")
	sort.Sort(ByUnits{products})
	fmt.Println(products)

	//Sort the products by id or name or cost or units or category
	//DONOT write your own implementation of sort
	//USE sort.Sort() function

	//Sort the product by id & print the list
	//Sort the products by name & print the list

}
