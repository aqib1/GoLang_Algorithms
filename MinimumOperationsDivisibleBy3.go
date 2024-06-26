package main

func main() {

}

func minimumOperations(nums []int) int {
	var count int = 0

	for i := 0; i < len(nums); i++ {
		if nums[i]%3 != 0 {
			count++
		}
	}

	return count
}
