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
	fmt.Println(name2)
}
