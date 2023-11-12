package dp

func LongestPalindrome(s string) string {
	dp := make([][]int, len(s))
	count := [2]int{0, 0}

	for i := 0; i < len(dp); i++ {
		dp[i] = append(dp[i], make([]int, len(s))...)
	}

	for i := range dp {
		dp[i][i] = 1
	}

	for j := 0; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if i == j-1 && s[i] == s[j] {
				dp[i][j] = 1
				if count[1]-count[0] < j-i {
					count[0] = i
					count[1] = j
				}
			} else if dp[i+1][j-1] == 1 && s[i] == s[j] {
				dp[i][j] = 1
				if count[1]-count[0] < j-i {
					count[0] = i
					count[1] = j
				}
			}
		}
	}

	return s[count[0] : count[1]+1]
}
