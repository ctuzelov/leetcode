package twoPointers

func IsPalindrome(s string) bool {
	t := make(map[string]int)
	flag := false
	p1 := 0
	var c rune
	p2 := len(t) - 1
	for p1 < len(t)/2 {
		c = rune(t[p1])
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= 48 && c <= 57) {
			_, ok := t[string(s[p1])]
			if ok {
				t[string(s[p1])] = +1
				flag = true
			} else {
				t[string(s[p1])] = 1
			}
		}
		c = rune(t[p2])
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= 48 && c <= 57) {
			_, ok := t[string(s[p1])]
			if ok {
				t[string(s[p2])] = -1
				flag = true
			} else {
				t[string(s[p1])] = 1
			}
		}
		if t[string(s[p1])] != 0 && flag {
			return false
		}

	}
	return true
}
