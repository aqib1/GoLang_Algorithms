package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Loop with slices
	var testSlice = []string{"Hi", "I", "am", "a", "slice"}

	for i := 0; i < len(testSlice); i++ {
		fmt.Print(testSlice[i] + " ")
	}
}
