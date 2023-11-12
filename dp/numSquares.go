package dp

import "fmt"

func NumSquares(n int) int {
	dp := make([]int, n+1)

	for i := 0; i < n+1; i++ {
		dp[i] = n + 1
	}

	dp[0] = 0

	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			n := j * j

			if i < n {
				continue
			}

			dp[i] = Min(dp[i], dp[i-n]+1)
		}
	}

	fmt.Println(dp)

	return dp[len(dp)-1]
}

/*
	12

	4 4 4

	1 -> 1
	2 -> 1 + 1
	3 -> 1 + 1 + 1
	4 -> 4
	5 -> 4 + 1
	6 -> 4 + 1 + 1
	7 -> 4 + 1 + 1 + 1
	8 -> 4 + 4
	9 -> 9
   10 -> 9 + 1
   11 -> 9 + 1 + 1
   12 -> 4 + 4 + 4
*/
