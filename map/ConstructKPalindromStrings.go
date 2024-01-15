package _map

// Time complexity O(n)
// Space complexity 0(26)
func canConstruct(s string, k int) bool {
	if len(s) < k {
		return false
	}

	var count = make([]int, 26)
	var oddCount = 0
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}

	for i := 0; i < len(count); i++ {
		if count[i] != 0 && count[i]%2 != 0 {
			oddCount++
		}
	}

	return oddCount <= k
}
