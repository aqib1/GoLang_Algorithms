package strings

// Time O(n)
// Space O(1)
func reverseString(s []byte) {
	start := 0
	end := len(s) - 1

	for start < end {
		tmp := s[start]
		s[start] = s[end]
		s[end] = tmp
		start++
		end--
	}
}
