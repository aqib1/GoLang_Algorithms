package main

import "fmt"

func main() {
	var slice = []int{}
	fmt.Println(slice)
	slice = append(slice, 1)
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))
}
