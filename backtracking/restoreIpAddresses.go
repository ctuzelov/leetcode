package backtracking

import (
	"strconv"
	"strings"
)

func RestoreIpAddresses(s string) []string {
	var (
		res       []string
		backtrack func(s string, curr []string)
	)

	backtrack = func(s string, curr []string) {
		if s == "" && len(curr) == 4 {
			res = append(res, strings.Join(curr, "."))
			return
		}

		if len(curr) > 4 {
			return
		}
		for i := 0; i < 3 && i < len(s); i++ {
			if isValid(s[0 : i+1]) {
				backtrack(s[i+1:], append(curr, s[0:i+1]))
			}
		}
	}

	backtrack(s, nil)
	return res
}

func isValid(s string) bool {
	n, _ := strconv.Atoi(s)

	if n > 255 {
		return false
	}

	if prefix := s[0]; len(s) != 1 && prefix == 48 {
		return false
	}
	return true
}
