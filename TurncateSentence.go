package main

import (
	"fmt"
	"strings"
)

func truncateSentence(s string, k int) string {
	if k == len(s) || k > len(s) {
		return s
	}

	var trimString = ""
	var words = strings.Split(s, " ")

	for i := 0; i < k; i++ {
		trimString += words[i] + " "
	}

	return strings.TrimSpace(trimString)
}

func main() {
	fmt.Println(truncateSentence("What is the solution to this problem", 4))
}
