package main

import "strings"
import "fmt"

func main() {
	fmt.Println(finalValueAfterOperations([]string{"--X", "X++", "X++"}))
}

func finalValueAfterOperations(operations []string) int {
	var result = 0
	var operators = []string{"++", "--"}

	for _, operation := range operations {
		if strings.HasPrefix(operation, operators[0]) ||
			strings.HasSuffix(operation, operators[0]) {
			result++
		} else if strings.HasPrefix(operation, operators[1]) ||
			strings.HasSuffix(operation, operators[1]) {
			result--
		}
	}
	return result
}
