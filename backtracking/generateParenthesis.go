package backtracking

func GenerateParenthesis(n int) []string {
	var (
		res       []string
		path      string
		backtrack func(left, right int)
	)

	backtrack = func(left, right int) {
		if left == 0 && right == 0 {
			cpy := make([]byte, len(path))
			copy(cpy, []byte(path))
			res = append(res, string(cpy))
			return
		}

		if left > 0 {
			path += string('(')
			backtrack(left-1, right)
			path = path[:len(path)-1]
		}
		if left < right {
			path += string(')')
			backtrack(left, right-1)
			path = path[:len(path)-1]
		}
	}

	backtrack(n, n)

	return res
}

/*

				((((
			(((
				((()
		((
				(()(
			(()
				(())
	(
				()((
			()(
				()()
		()
				())(
			())
				()))


			)()
		)(
	)
		))
			)))



*/
