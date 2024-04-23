package main

import "fmt"

func main() {
	var message string = "Hello World"              // type declaration
	var inferredMessage = "I am inferred string"    // type inferred declaration
	anotherInferredMessage := "Hi i am also string" // type inferred declaration which can be declared only inside
	// function and declaration should be on same line

	fmt.Println(message)
	fmt.Println(inferredMessage)
	fmt.Println(anotherInferredMessage)

	var a1 string              // Declaration
	a1 = "I am another string" // initialization

	var c1, c2, c3, c4 int = 1, 2, 3, 4

	fmt.Println(a1)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
}
