package main

import "math"

func findPermutationDifference(s string, t string) int {
	sCharsIndexMap := make(map[string]int)
	var difference float64 = 0

	for i, r := range s {
		sCharsIndexMap[string(r)] = i
	}

	for i, r := range t {
		difference += math.Abs(float64(i) - float64(sCharsIndexMap[string(r)]))
	}

	return int(difference)
}
