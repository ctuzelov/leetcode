package backtracking

func SolveNQueens(n int) [][]string {
	var (
		res       [][]string
		board     [][]byte
		backtrack func(row int)
		isValid   func(row int, col int) bool
	)

	for i := 0; i < n; i++ {
		temp := make([]byte, n)
		for j := 0; j < n; j++ {
			temp[j] = '.'
		}
		board = append(board, temp)
	}

	isValid = func(row, col int) bool {
		for i := 0; i < row; i++ {
			if board[i][col] == 'Q' {
				return false
			}
		}

		for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
			if board[i][j] == 'Q' {
				return false
			}
		}

		for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
			if board[i][j] == 'Q' {
				return false
			}
		}

		return true
	}

	backtrack = func(row int) {
		if row == n {
			var temp []string
			for i := 0; i < n; i++ {
				temp = append(temp, string(board[i]))
			}
			res = append(res, temp)
		}

		for col := 0; col < n; col++ {
			if isValid(row, col) {
				board[row][col] = 'Q'
				backtrack(row + 1)
				board[row][col] = '.'
			}
		}
	}

	backtrack(0)

	return res
}
