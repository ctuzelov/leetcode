package dp

func IntegerBreak(n int) int {
	if n == 2 || n == 3 {
		return 1
	}

	dp := make([]int, n+1)

	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j < i; j++ {
			temp := dp[j] * dp[i-j]
			dp[i] = max(dp[i], temp)
		}
	}

	return dp[len(dp)-1]
}
