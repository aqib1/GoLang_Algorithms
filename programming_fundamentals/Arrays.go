package main

import "fmt"

func main() {
	var arr1 = [2]int{1, 2} // define an array with len

	fmt.Println(len(arr1))
	fmt.Println(arr1)

	var arr2 = [...]int{1, 2, 3, 4, 5} // define an array with inferred size
	fmt.Println(len(arr2))
	fmt.Println(arr2)

	var arr3 = [3]int{} // define an array without initialization of values
	arr3[0] = 1
	arr3[1] = 2
	arr3[2] = 3

	fmt.Println(arr3)

	var arr4 = [5]int{1: 5, 4: 3} // what if we only want to assign value to index 2, and 4
	// 1 index is assigned with 5 and 4rth with 3 rest will be zero
	fmt.Println(arr4)
}
