package main

import (
	"fmt"
	"math"
	"strings"
)

func mostWordsFound(sentences []string) int {
	var result = math.MinInt64
	for _, sentence := range sentences {
		var wordsLen = len(strings.Split(sentence, " "))
		if result < wordsLen {
			result = wordsLen
		}
	}

	return result
}

func main() {
	fmt.Println(mostWordsFound([]string{
		"alice and bob love leetcode",
		"i think so too",
		"this is great thanks very much",
	}))
}
