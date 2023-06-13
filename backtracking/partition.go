package backtracking

func Partition(s string) [][]string {
	var (
		res       [][]string
		backtrack func(state []string, s string)
	)

	backtrack = func(state []string, s string) {
		if s == "" {
			cpy := make([]string, len(state))
			copy(cpy, state)
			res = append(res, cpy)
			return
		}

		for i := 0; i < len(s); i++ {
			if isPalindrom(s[:i+1], 0, len(s[:i+1])-1) {
				state = append(state, s[:i+1])
				backtrack(state, s[i+1:])
				state = state[:len(state)-1]
			}
		}
	}

	backtrack([]string{}, s)
	return res
}

func isPalindrom(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
