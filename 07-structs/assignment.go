package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
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
	costlyProductPredicate := func(product Product) bool {
		return product.Cost > 100
	}
	costlyProducts := Filter(products, costlyProductPredicate)
	fmt.Println("Costly Products")
	Print(costlyProducts)

	stationaryProductPredicate := func(product Product) bool {
		return product.Category == "Stationary"
	}
	stationaryProducts := Filter(products, stationaryProductPredicate)
	fmt.Println("Stationary Products")
	Print(stationaryProducts)

	fmt.Println("Are there any costly products ? : ", Any(products, costlyProductPredicate))

	fmt.Println("Are there any stationary products ? : ", Any(products, stationaryProductPredicate))

	fmt.Println("Are all the products stationary products ? : ", All(products, stationaryProductPredicate))

	fmt.Println("Are all the products costly products ? : ", All(products, costlyProductPredicate))
}

func Print(products []Product) {
	for _, p := range products {
		fmt.Printf("%+v\n", p)
	}
}

//IndexOf => returns the index of the given product in the products collections (-1 if it doesnt exist)

func IndexOf(products []Product, product Product) int {
	for i, p := range products {
		if p == product {
			return i
		}
	}
	return -1
}

//Includes => return true/false based on the existence of the product in the products collection
func Includes(products []Product, product Product) bool {
	return IndexOf(products, product) != -1
}

//Filter => returns a new collection of products that match the given criteria
//use cases
//filter costly products (cost > 100)
/*
func FilterCostlyProducts(products []Product) []Product {
 	var costlyProducts []Product
	for _, p := range products {
		if p.Cost > 100 {
			costlyProducts = append(costlyProducts, p)
		}
	}
	return costlyProducts
} */

//filter stationary products (category = "Stationary")
/* func FilterStationaryProducts(products []Product) []Product {
	var stationaryProducts []Product
	for _, p := range products {
		if p.Category == "Stationary" {
			stationaryProducts = append(stationaryProducts, p)
		}
	}
	return stationaryProducts
} */

func Filter(products []Product, predicate func(Product) bool) []Product {
	var result []Product
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
func Any(products []Product, predicate func(Product) bool) bool {
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
func All(products []Product, predicate func(Product) bool) bool {
	for _, p := range products {
		if !predicate(p) {
			return false
		}
	}
	return true
}
