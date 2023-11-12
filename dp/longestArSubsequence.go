package dp

func LongestSubsequence(arr []int, difference int) int {
	dp := make(map[int]int)
	mxLen := 0

	for _, n := range arr {
		if _, ok := dp[n-difference]; ok {
			dp[n] = dp[n-difference] + 1
			mxLen = Max(mxLen, dp[n])
		} else {
			dp[n] = 1
		}
	}
	/*
		1, 7, 8, 5, 6, 3, 4, 2, 1, 0
	*/

	return mxLen
}
