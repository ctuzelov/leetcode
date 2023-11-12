package dp

func MinCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))

	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < len(cost); i++ {
		dp[i] = cost[i] + Min(dp[i-1], dp[i-2])
	}

	return Min(dp[len(dp)-1], dp[len(dp)-2])
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
0,1,100,1,1,1,100,1,1,30,1

*/
