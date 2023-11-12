package dp

import "fmt"

func LongestCommonSubsequence(text2 string, text1 string) int {
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = append(dp[i], make([]int, len(text2)+1)...)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	fmt.Println(dp)

	return dp[len(text1)][len(text2)]
}

/*

text1 = "abcde", text2 = "ace"

[a b c d e]
[a c e]

dp = [0,0,0,0,0,0]

a =


*/
