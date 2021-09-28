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

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type By func(p1, p2 *Product) bool

func (by By) Sort(products []Product) {
	ps := &productSorter{
		products: products,
		by:       by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}

// Len is part of sort.Interface.
func (s *productSorter) Len() int {
	return len(s.products)
}

// Swap is part of sort.Interface.
func (s *productSorter) Swap(i, j int) {
	s.products[i], s.products[j] = s.products[j], s.products[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *productSorter) Less(i, j int) bool {
	return s.by(&s.products[i], &s.products[j])
}

// productSorter joins a By function and a slice of Products to be sorted.
type productSorter struct {
	products []Product
	by       func(p1, p2 *Product) bool // Closure used in the Less method.
}

func main() {
	products := []Product{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	id := func(p1, p2 *Product) bool {
		return p1.Id < p2.Id
	}
	name := func(p1, p2 *Product) bool {
		return p1.Name < p2.Name
	}
	cost := func(p1, p2 *Product) bool {
		return p1.Cost < p2.Cost
	}
	units := func(p1, p2 *Product) bool {
		return p1.Units < p2.Units
	}
	category := func(p1, p2 *Product) bool {
		return p1.Category < p2.Category
	}

	By(id).Sort(products)
	fmt.Println("By Id:", products)

	By(name).Sort(products)
	fmt.Println("By name:", products)

	By(cost).Sort(products)
	fmt.Println("By cost:", products)

	By(units).Sort(products)
	fmt.Println("By units:", products)

	By(category).Sort(products)
	fmt.Println("By category:", products)

	//Sort the products by id or name or cost or units or category
	//DONOT write your own implementation of sort
	//USE sort.Sort() function

	//Sort the product by id & print the list
	//Sort the products by name & print the list

}
