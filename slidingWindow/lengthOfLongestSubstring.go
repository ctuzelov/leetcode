package slidingWindow

func LengthOfLongestSubstring(s string) int {
	var (
		res   int
		start int
		m     = make(map[byte]int)
		max   func(a, b int) int
	)

	max = func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	for end := 0; end < len(s); end++ {
		m[s[end]]++

		for m[s[end]] > 1 {
			m[s[start]]--
			start++
		}
		res = max(res, end-start+1)
	}
	return res
}
