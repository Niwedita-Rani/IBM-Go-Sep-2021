package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

//Creating an ALIAS for a slice so that we can have methods on it
type Products []Product

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
	fmt.Println(products.Format())

	fmt.Println("Filtered List")
	costlyProductPredicate := func(product Product) bool {
		return product.Cost > 100
	}

	costlyProducts := products.Filter(costlyProductPredicate)
	fmt.Println("Costly Products")
	fmt.Println(costlyProducts.Format())

	stationaryProductPredicate := func(product Product) bool {
		return product.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductPredicate)
	fmt.Println("Stationary Products")
	fmt.Println(stationaryProducts.Format())

	fmt.Println("Any?")
	fmt.Println("Are there any costly products ? : ", products.Any(costlyProductPredicate))

	fmt.Println("Are there any stationary products ? : ", products.Any(stationaryProductPredicate))

	fmt.Println("All?")
	fmt.Println("Are all the products stationary products ? : ", products.All(stationaryProductPredicate))

	fmt.Println("Are all the products costly products ? : ", products.All(costlyProductPredicate))
}

func (product Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s\n", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func (products Products) Format() string {
	result := ""
	for _, p := range products {
		result += fmt.Sprintf("%s", p.Format())
	}
	return result
}

//IndexOf => returns the index of the given product in the products collections (-1 if it doesnt exist)

func (products Products) IndexOf(product Product) int {
	for i, p := range products {
		if p == product {
			return i
		}
	}
	return -1
}

//Includes => return true/false based on the existence of the product in the products collection
func (products Products) Includes(product Product) bool {
	return products.IndexOf(product) != -1
}

func (products Products) Filter(predicate func(Product) bool) Products {
	var result Products
	for _, p := range products {
		if predicate(p) {
			result = append(result, p)
		}
	}
	return result
}

//Any => return true/false based on the existence of the given product that satisfies the given criteria in the products collection
//use cases
//Are there any costly products?
//Are there any stationary products?
func (products Products) Any(predicate func(Product) bool) bool {
	for _, p := range products {
		if predicate(p) {
			return true
		}
	}
	return false
}

//All => return true/false based on the existence of all the products that satisfy the given criteria in the products collection
//use cases
//Are all the products stationary products?
//Are all the products costly products?
func (products Products) All(predicate func(Product) bool) bool {
	for _, p := range products {
		if !predicate(p) {
			return false
		}
	}
	return true
}
