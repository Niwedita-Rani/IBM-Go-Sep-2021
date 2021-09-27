package main

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
}

//IndexOf => returns the index of the given product in the products collections (-1 if it doesnt exist)

//Includes => return true/false based on the existence of the product in the products collection

//Filter => returns a new collection of products that match the given criteria
//use cases
//filter costly products (cost > 100)
//filter stationary products (category = "Stationary")

//Any => return true/false based on the existence of the given product that satisfies the given criteria in the products collection
//use cases
//Are there any costly products?
//Are there any stationary products?

//All => return true/false based on the existence of all the products that satisfy the given criteria in the products collection
//use cases
//Are all the products stationary products?
//Are all the products costly products?
