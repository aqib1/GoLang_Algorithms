package main

import "math"

func main() {

}

func maximumWealth(accounts [][]int) int {

	var maximum int = math.MinInt64

	for i := 0; i < len(accounts); i++ {
		var count int = 0
		for j := 0; j < len(accounts[i]); j++ {
			count += accounts[i][j]
		}
		if count > maximum {
			maximum = count
		}
	}

	return maximum
}
