package dp

func MinDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word2)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = append(dp[i], make([]int, len(word1)+1)...)
	}

	for j := range word1 {
		dp[len(word2)][j] = len(word1) - j
	}
	for i := range word2 {
		dp[i][len(word1)] = len(word2) - i
	}

	for i := len(word2) - 1; i >= 0; i-- {
		for j := len(word1) - 1; j >= 0; j-- {
			if word1[j] == word2[i] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = 1 + min(dp[i+1][j+1], dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[0][0]
}
