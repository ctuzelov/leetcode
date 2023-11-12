package dp

func NumberOfWays(startPos int, endPos int, k int) int {
	var dfs func(stepNumbers, step int) int
	m := make(map[[2]int]int)

	dfs = func(stepNumbers, step int) int {
		if stepNumbers < 0 {
			return 0
		}

		if stepNumbers == 0 && step == endPos {
			return 1
		}

		if val, ok := m[[2]int{stepNumbers, step}]; ok {
			return val
		}

		res := dfs(stepNumbers-1, step-1) + dfs(stepNumbers-1, step+1)
		m[[2]int{stepNumbers, step}] = res
		return res
	}

	return dfs(k, startPos)
}

/*

var dfs func(i, k int) int

	dfs = func(i, k int) int {
		// Check if k is negative or if i != j and k is 0
		if k < 0 {
			return 0
		}

		// Check if we have reached the destination and k is 0
		if i == endPos && k == 0 {
			return 1
		}

		// Calculate the result using recursive calls
		result := (dfs(i+1, k-1) + dfs(i-1, k-1))

		return result
	}

	return dfs(startPos, k)

*/
