package main

import "fmt"

func createAllSubstrings(str string) []string {
	var substrings []string
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			substrings = append(substrings, str[i:j])
		}
	}
	return substrings
}

func balancedStringSplit(s string) int {
	var count = 0
	var balanced = 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'R' {
			balanced++
		}
		if s[i] == 'L' {
			balanced--
		}

		if balanced == 0 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(balancedStringSplit("LLLLRRRR"))
}
