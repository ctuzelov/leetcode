package backtracking

func Exist(board [][]byte, word string) bool {
	var (
		res       = 0
		backtrack func(x int, y int, valid int, board [][]byte)
	)

	backtrack = func(x int, y int, valid int, board [][]byte) {
		if x < 0 || y < 0 || x > len(board)-1 || y > len(board[0])-1 || board[x][y] != word[valid-1] || board[x][y] == '*' {
			return
		}

		original := board[x][y]
		board[x][y] = '*'

		if valid == len(word) {
			res = valid
			return
		}
		backtrack(x, y+1, valid+1, board)
		backtrack(x+1, y, valid+1, board)
		backtrack(x, y-1, valid+1, board)
		backtrack(x-1, y, valid+1, board)
		board[x][y] = original

	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			backtrack(i, j, 1, board)
			if res == len(word) {
				return true
			}
			res = 0
		}
	}
	return false
}
