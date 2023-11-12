package dp

func NumSplits(s string) int {

	dp := make([]int, len(s))
	m := make(map[byte]int)
	dp[0] = 1
	m[s[0]] = 1

	for i := 1; i < len(s); i++ {
		c := s[i]
		m[c] += 1
	}

	count := 0

	ml := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		c := s[i]
		ml[c] += 1
		m[c] -= 1

		if m[c] == 0 {
			delete(m, c)
		}

		if len(ml) == len(m) {
			count += 1
		}
	}
	// aacaba

	return count
}

/*



 */
