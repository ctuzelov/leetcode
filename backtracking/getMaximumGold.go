package backtracking

func GetMaximumGold(grid [][]int) int {
	var (
		res       int
		temp      int
		backtrack func(x, y int, max int, grid [][]int)
	)

	backtrack = func(x, y, max int, grid [][]int) {
		if x < 0 || y < 0 || x > len(grid)-1 || y > len(grid[0])-1 || grid[x][y] == -1 || grid[x][y] == 0 {
			return
		}

		temp += grid[x][y]
		if res < temp {
			res = temp
		}

		origin := grid[x][y]
		grid[x][y] = -1

		backtrack(x, y+1, temp, grid)
		backtrack(x+1, y, temp, grid)
		backtrack(x, y-1, temp, grid)
		backtrack(x-1, y, temp, grid)

		grid[x][y] = origin
		temp -= grid[x][y]
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			backtrack(i, j, temp+grid[i][j], grid)
			temp = 0
		}
	}

	return res
}

/*
{1,0,7},
{2,0,6},
{3,4,5},
{0,3,0},
{9,0,20}

*/
