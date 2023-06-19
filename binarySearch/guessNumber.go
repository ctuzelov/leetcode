package binarySearch

func GuessNumber(n int) int {
	left, right := 0, n+1
	for right-left > 1 {
		mid := (left + right) >> 1
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == 1 {
			left = mid
		} else if guess(mid) == -1 {
			right = mid
		}
	}
	return left
}

func guess(n int) int {
	return 0
}
