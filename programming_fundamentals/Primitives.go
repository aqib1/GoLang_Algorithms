package main

import "fmt"

func main() {
	var name string = "Aqib"
	// := can not be used outside of function
	name1 := "Aqib"
	var name2 string

	fmt.Println(name != name1)
	fmt.Println(name2)
	name2 = "Maria"
	numberFloat := 37863.5557527
	fmt.Printf("My name is %T\n", name2)
	fmt.Printf("you scored %0.2f\n", numberFloat)

	var formatedStr = fmt.Sprintf("you scored %0.2f\n", numberFloat)
	fmt.Println("Saved string is = ", formatedStr)
}
