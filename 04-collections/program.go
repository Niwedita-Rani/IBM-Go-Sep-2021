package main

import "fmt"

func main() {
	//Arrays
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	var nos [5]int = [...]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	//Iterating an array (using the indexer)
	for idx := 0; idx < len(nos); idx++ {
		fmt.Println(nos[idx])
	}

	//Iterating an array (using range)
	for _, val := range nos {
		fmt.Println(val)
	}

	//creating a copy of an array
	var nos2 [5]int = nos
	fmt.Println(nos2)
	nos2[0] = 200

	fmt.Println(nos, nos2)

	//Slice
	/*
		var names []string
		printStats("Empty slice", names)
	*/
	//proactively allocating memory
	var names []string = make([]string, 10, 10)

	//intializing the slice values when creating it
	//var names []string = []string{"Magesh", "Suresh", "Ramesh"}
	names = append(names, "Magesh", "Suresh", "Ramesh")

	fmt.Println(names)
	printStats("slice intialized with values", names)
	printAddresses(names)

	//adding elements to a slice
	//names = append(names, "Rajesh")
	//names = append(names, "Rajesh", "Ganesh")
	//newNames := []string{"Rajesh", "Ganesh"}
	names = append(names, "Rajesh")
	printStats("after appending 1 value", names)
	printAddresses(names)

	//names = append(names, newNames...)
	names = append(names, "Ganesh")
	names = append(names, "John")
	printStats("after appending values to the capacity", names)

	names = append(names, "Philip")
	printStats("after appending values beyond capacity", names)
	fmt.Println(names)

	/*

	 */

	//Iterating a slice
	for idx, val := range names {
		fmt.Println(idx, val)
	}

	//Slicing a slice
	/*
		slice[lo:hi] => from index lo to hi-1
		slice[lo: ] => from index lo to end
		slice[: hi] => from start to hi-1
		slice[:] => from start to end
		slice[lo : lo] => []
		slice[lo : lo+1] => slice[lo]
	*/
	/*
		fmt.Println("names[1:3] => ", names[1:3])
		fmt.Println("names[:3] => ", names[:3])
		fmt.Println("names[3: ] => ", names[3:])
	*/

	//Creating a slice from an array
	var newNos [5]int = [...]int{3, 1, 4, 2, 5}
	var nosSlice = newNos[:]
	fmt.Printf("%T, %T\n", newNos, nosSlice)
}

func printStats(label string, data []string) {
	fmt.Printf("[%s] Len = %d, Cap = %d\n", label, len(data), cap(data))
}

func printAddresses(data []string) {
	fmt.Println()
	fmt.Println("element addresses:")
	for idx := range data {
		fmt.Printf("%d: %p\n", idx, &data[idx])
	}
	fmt.Println()
}
