package November

func IsReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	diag := min(abs(fy-sy), abs(fx-sx))
	addit := max(abs(fy-sy), abs(fx-sx)) - diag

	if diag == 0 && addit == 0 && t <= 1 || diag+addit < t {
		return false
	}

	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*

sx = 2, sy = 4, fx = 7, fy = 7, t = 6


*/
