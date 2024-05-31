package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(scoreOfString("hello"))
}

func scoreOfString(s string) int {
	charArray := make([]int32, len(s))
	var sum float64 = 0

	for i, r := range s {
		charArray[i] = r
	}

	for i := 0; i < len(charArray)-1; i++ {
		sum += math.Abs(float64(charArray[i]) - float64(charArray[i+1]))
	}

	return int(sum)
}
