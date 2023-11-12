package dp

func CountSubstrings(s string) int {
	dp := make([][]int, len(s))
	count := 0

	for i := 0; i < len(dp); i++ {
		dp[i] = append(dp[i], make([]int, len(s))...)
	}

	for i := range dp {
		dp[i][i] = 1
		count++
	}

	for j := 1; j < len(dp); j++ {
		for i := 0; i < j; i++ {
			if i == j-1 && s[i] == s[j] {
				dp[i][j] = 1
				count++
			} else if dp[i+1][j-1] == 1 && s[i] == s[j] {
				dp[i][j] = 1
				count++
			}
		}
	}
	return count
}
