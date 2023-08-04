package backtracking

import "strings"

func SolveNQueens(n int) [][]string {
	board := make([][]rune, n)
	for r := 0; r < n; r++ {
		board[r] = make([]rune, n)
		for c := 0; c < n; c++ {
			board[r][c] = '.'
		}
	}

	res := [][]string{}
	placeQueen(0, n, board, &res)
	return res
}

func placeQueen(qId, n int, board [][]rune, res *[][]string) {
	if qId == n {
		addToResult(board, res)
		return
	}

	for row := 0; row < n; row++ {
		if isValidCell(board, row, qId) {
			board[row][qId] = 'Q'
			placeQueen(qId+1, n, board, res)
			board[row][qId] = '.'
		}
	}
}

func isValidCell(board [][]rune, row, col int) bool {
	n := len(board)

	// Same Row
	for c := 0; c < col; c++ {
		if board[row][c] == 'Q' {
			return false
		}
	}

	// Top Left Diagonal
	for r, c := row-1, col-1; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if board[r][c] == 'Q' {
			return false
		}
	}

	// Bottom Left Diagonal
	for r, c := row+1, col-1; r < n && c >= 0; r, c = r+1, c-1 {
		if board[r][c] == 'Q' {
			return false
		}
	}

	return true
}

func addToResult(board [][]rune, res *[][]string) {
	n := len(board)
	rows := []string{}

	for r := 0; r < n; r++ {
		row := strings.Builder{}
		for c := 0; c < n; c++ {
			row.WriteRune(board[r][c])
		}
		rows = append(rows, row.String())
	}

	*res = append(*res, rows)
}
