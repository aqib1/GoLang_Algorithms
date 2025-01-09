package main

import "fmt"

func main() {
	fmt.Println(twoSum2([]int{2, 7, 11, 15}, 9))
}

// Time complexity O(n)2 & Space complexity O(2)
func twoSum(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = []int{i, j}
				break
			}
		}
	}

	return result
}

// Time complexity O(n) and space O(n)
func twoSum2(nums []int, target int) []int {
	var result []int
	var numberByIndex = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		value, ok := numberByIndex[target-nums[i]]
		if ok {
			result = []int{value, i}
		} else {
			numberByIndex[nums[i]] = i
		}
	}

	return result
}
