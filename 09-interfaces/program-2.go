package main

import "fmt"

func main() {
	var x interface{}
	//x = 100

	//x = "hello"
	//x = true
	// x = struct{}{}
	//x = func() {}
	x = []int{4, 1, 3, 2, 5}

	typ := fmt.Sprintf("%T", x)
	fmt.Println(typ)

	if no, ok := x.(int); ok {
		fmt.Println(no + 200)
	} else {
		fmt.Println("not an int")
	}

	switch v := x.(type) {
	case int:
		fmt.Println(v + 200)
	case string:
		fmt.Println(v + " world")
	case bool:
		fmt.Println("boolean value => ", v)
	case []int:
		fmt.Println("array => ", v)
	default:
		fmt.Println("unknown type")

	}
}
