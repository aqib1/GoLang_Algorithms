package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(largestLocal(
		[][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}))
}

func largestLocal(grid [][]int) [][]int {
	var result = make([][]int, len(grid[0])-2, len(grid)-2)

	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[i])-2; j++ {
			result[i] = append(result[i], getMaxNumber(grid, i, j))

		}
	}

	return result
}

func getMaxNumber(grid [][]int, i int, j int) int {
	var maxValue = math.MinInt64

	for x := i; x < i+3; x++ {
		for y := j; y < j+3; y++ {
			maxValue = int(math.Max(float64(maxValue), float64(grid[x][y])))
		}
	}

	return maxValue
}
