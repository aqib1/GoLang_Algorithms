package main

import "fmt"

const PI float64 = 3.1415

const (
	NAME     = "AQIB"
	LASTNAME = "JAVED"
)

func main() {
	fmt.Println(PI)
	fmt.Print(NAME, "\n")
	fmt.Printf("%v\n", LASTNAME) // %v is used to print value of LASTNAME which is JAVED
	fmt.Printf("%T", LASTNAME)   // %T is used to print type of LASTNAME which is string
}
