package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(10, 20))                                                 //=> 30
	fmt.Println(sum(10, 20, 30, 40))                                         //=> 100
	fmt.Println(sum(10))                                                     //=> 10
	fmt.Println(sum())                                                       //=> 0
	fmt.Println(sum(10, "20", 30, "40"))                                     //=> 100
	fmt.Println(sum(10, "20", 30, "40", "abc"))                              //=> 100
	fmt.Println(sum(10, 20, []int{30, 40}))                                  //=> 100
	fmt.Println(sum(10, 20, []interface{}{30, 40, []int{10, 20}}))           //=> 130
	fmt.Println(sum(10, 20, []interface{}{30, 40, []interface{}{10, "20"}})) //=> 130
}

func sum(data ...interface{}) int {
	result := 0
	for _, v := range data {
		switch value := v.(type) {
		case int:
			result += value
		case string:
			if i, err := strconv.Atoi(value); err == nil {
				result += i
			}
		case []int:
			intfList := make([]interface{}, len(value))
			for i, vInt := range value {
				intfList[i] = vInt
			}
			result += sum(intfList...)
		case []interface{}:
			result += sum(value...)
		}
	}
	return result
}
