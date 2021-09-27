package main

import "fmt"

func main() {
	var x int = 10
	var xPtr *int = &x
	fmt.Println(xPtr, x)

	//deferencing a pointer (accessing the underlying value using the pointer)
	var y int = *xPtr
	fmt.Println(x, y)

	fmt.Println("Before incrementing x = ", x)
	increment(&x)
	fmt.Println("After incrementing x = ", x)

	a, b := 10, 20
	fmt.Println("Before swapping a = ", a, "b = ", b)
	swap( /* a, b */ )
	fmt.Println("After swapping a = ", a, "b = ", b)
}

func increment(no *int) {
	*no++
}

func swap( /*  */ ) {

}
