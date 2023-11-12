package dp

func Change(coins []int, amount int) int {
	dp := make([]int, amount+1)

	dp[0] = 1

	for _, c := range coins {
		for a := c; a <= amount; a++ {
			dp[a] += dp[a-c]
		}
	}

	return dp[len(dp)-1]
}
