package backtracking

func LetterCombinations(digits string) []string {
	var (
		res       []string
		backtrack func(num int, path string)
		alph      = []string{"", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	)

	if len(digits) == 0 {
		return []string{}
	}

	backtrack = func(num int, path string) {
		if len(path) == len(digits) {
			cp := make([]byte, len(path))
			copy(cp, []byte(path))
			res = append(res, string(cp))
			return
		}

		for i := num; i < len(digits); i++ {
			for j := 0; j < len(alph[digits[i]-49]); j++ {
				path += string(alph[digits[i]-49][j])
				backtrack(i+1, path)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0, "")
	return res
}
