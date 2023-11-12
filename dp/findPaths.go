package dp


func FindPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	var dfs func(int, int, int) int
	var dp = make([][][]int, m)
	var mod int = 1e9 + 7

	for i := 0; i < m; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, maxMove+1)
			for k := 0; k < maxMove+1; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	dfs = func(curr_move, row, column int) int {
		if row >= m || column >= n || row < 0 || column < 0 {
			return 1
		}

		if curr_move == 0 {
			return 0
		}

		if dp[row][column][curr_move] != -1 {
			return dp[row][column][curr_move]
		}

		res := dfs(curr_move-1, row, column-1) + dfs(curr_move-1, row, column+1) + dfs(curr_move-1, row+1, column) + dfs(curr_move-1, row-1, column)
		dp[row][column][curr_move] = res % mod
		return res % mod
	}

	r := dfs(maxMove, startRow, startColumn)
	return r

}
