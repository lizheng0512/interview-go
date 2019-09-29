package string

func LengthOfLongestSubstring(s string) int {
	subString := make(map[string]int)
	start := -1
	maxLength := 0
	for i, c := range s {
		char := string(c)
		if n, ok := subString[char]; ok {
			if n > start {
				start = n
			}
		}
		subString[char] = i
		if i-start > maxLength {
			maxLength = i - start
		}
	}
	return maxLength
}
