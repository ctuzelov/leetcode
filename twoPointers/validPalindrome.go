package twoPointers

func ValidPalindrome(s string) bool {
	var valid func(i, j, count int) bool

	valid = func(i, j, count int) bool {
		if count > 1 {
			return false
		}
		for i < j {
			if s[i] != s[j] {
				return valid(i+1, j, count+1) || valid(i, j-1, count+1)
			}
			i++
			j--
		}
		return true
	}

	return valid(0, len(s)-1, 0)
}
