package main

import "fmt"

func main() {
	fmt.Println("Welcome to the shuffle the array")
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}))
}

func shuffle(array []int) []int {
	midIndex := len(array) / 2
	result := make([]int, len(array))

	var pointer = 0
	for i, j := 0, midIndex; i <= midIndex && j < len(array); i, j = i+1, j+1 {
		result[pointer], result[pointer+1] = array[i], array[j]
		pointer += 2
	}

	return result
}
