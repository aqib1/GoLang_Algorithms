package main

import "strings"

import "fmt"

func main() {
	fmt.Println(findWordsContaining([]string{
		"leet", "code",
	}, 'e'))
}

func findWordsContaining(words []string, x byte) []int {
	var result = []int{}
	for i := 0; i < len(words); i++ {
		if strings.Contains(words[i], string(x)) {
			result = append(result, i)
		}
	}

	return result
}
