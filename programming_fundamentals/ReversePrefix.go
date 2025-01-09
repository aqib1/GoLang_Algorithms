package main

import (
	"fmt"
	"strings"
)

func reversePrefix(word string, ch byte) string {
	var chIndex = strings.IndexByte(word, ch)
	if chIndex == -1 {
		return word
	}
	return reverseString(word[:chIndex+1]) + word[chIndex+1:]
}

func reverseString(str string) string {
	var reversedStr = ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	return reversedStr
}

func main() {
	fmt.Println(reversePrefix("abcd", 'z'))
}
