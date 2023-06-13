package backtracking

func LetterCasePermutation(str string) []string {
	var (
		res       []string
		s         = []byte(str)
		backtrack func(ind int)
		letters   []int
	)

	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' || s[i] >= 'a' && s[i] <= 'z' {
			letters = append(letters, i)
		}
	}

	backtrack = func(ind int) {
		if ind == len(letters) {
			cpy := make([]byte, len(s))
			copy(cpy, s)
			res = append(res, string(s))
			return
		}

		for i := 0; i < 2; i++ {
			orig := s[letters[ind]]
			if i == 0 {
				s[letters[ind]] = toLower(s[letters[ind]])
			} else {
				s[letters[ind]] = toUpper(s[letters[ind]])
			}
			backtrack(ind + 1)
			s[letters[ind]] = orig
		}
	}
	backtrack(0)
	return res
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	} else {
		return b
	}
}

func toUpper(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - 32
	} else {
		return b
	}
}
