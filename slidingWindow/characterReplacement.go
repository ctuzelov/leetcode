package slidingWindow

func CharacterReplacement(s string, k int) int {
	var (
		start   int
		res     int
		freqArr [26]int
		freqMax int
		max     func(a, b int) int
	)

	max = func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	for end := 0; end < len(s); end++ {
		freqArr[s[end]-'A']++
		freqMax = max(freqMax, freqArr[s[end]-'A'])

		for end-start+1-freqMax > k {
			freqArr[s[start]-'A']--
			start++
		}

		res = max(res, end-start+1)
	}
	return res
}

/*
	AABABBA
*/
