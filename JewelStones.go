package main

func numJewelsInStones(jewels string, stones string) int {
	stonesCountMap := make(map[string]int)
	for _, stone := range stones {
		stonesCountMap[string(stone)] = stonesCountMap[string(stone)] + 1
	}

	count := 0

	for _, jewel := range jewels {
		count += stonesCountMap[string(jewel)]
	}

	return count
}
