package binarySearch

func Search(n []int, target int) int {
	left, right := -1, len(n)
	for right-left > 1 {
		mid := (left + right) >> 1
		if n[mid] == target {
			return mid
		} else if n[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return -1
}
