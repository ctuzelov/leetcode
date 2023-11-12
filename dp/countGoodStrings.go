package dp

var mod int = 1e9 + 7

func CountGoodStrings(low int, high int, zero int, one int) int {
	var dfs func(length int) int
	m := make(map[int]int)

	dfs = func(length int) int {
		if length > high {
			return 0
		}

		if n, ok := m[length]; ok {
			return n
		}

		if length <= high && length >= low {
			m[length] = (dfs(length+zero) + dfs(length+one) + 1) % mod
			return m[length]
		}

		m[length] = (dfs(length+zero) + dfs(length+one)) % mod
		return m[length]
	}

	return dfs(0)
}
