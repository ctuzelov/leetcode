package backtracking

func SolveSudoku(board [][]byte) {
	var (
		isValid   func(row, col int, element byte) bool
		backtrack func(row, col int) bool
	)

	isValid = func(row, col int, element byte) bool {
		for i := 0; i < 9; i++ {
			if board[i][col] == element {
				return false
			}
		}

		for j := 0; j < 9; j++ {
			if board[row][j] == element {
				return false
			}
		}

		for i := (row / 3) * 3; i < (row/3)*3+3; i++ {
			for j := (col / 3) * 3; j < (col/3)*3+3; j++ {
				if board[i][j] == element {
					return false
				}
			}
		}
		return true
	}

	backtrack = func(row, col int) bool {

		for i := row; i <= 8; i++ {
			for j := col; j <= 8; j++ {
				if board[i][j] != '.' {
					continue
				}
				for input := '1'; input <= '9'; input++ {
					if isValid(i, j, byte(input)) {
						board[i][j] = byte(input)
						if backtrack(i, j) {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
		return true
	}

	backtrack(0, 0)
}

/*

	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	{'.', '.', '.', '.', '8', '.', '.', '7', '9'}

*/
