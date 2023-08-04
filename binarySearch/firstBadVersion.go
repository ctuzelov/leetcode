package binarySearch

func FirstBadVersion(n int) int {
	l, r := 0, n+1
	for r-l > 1 {
		m := (l + r) / 2
		if isBadVersion(m) {
			r = m
		} else {
			l = m
		}
	}
	return r
}

func isBadVersion(n int) bool {
	return true
}
