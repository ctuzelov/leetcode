package twoPointers

func IsPalindrome(s string) bool {
	t := ""

	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			if c >= 'A' && c <= 'Z' {
				t += string(c + 32)
			} else {
				t += string(c)
			}
		}
		if c >= 48 && c <= 57 {
			t += string(c)
		}
	}
	p1 := 0
	p2 := len(t) - 1
	for p1 < len(t)/2 {
		if t[p1] != t[p2] {
			return false
		}
		p1++
		p2--
	}
	return true
}
