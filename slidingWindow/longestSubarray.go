package slidingWindow

func LongestSubarray(nums []int) int {
	var (
		m = map[int]int{
			0: 0,
			1: 0,
		}
		res   int
		start int
		max   func(a, b int) int
	)

	max = func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	for end := 0; end < len(nums); end++ {
		m[nums[end]]++
		for m[0] > 1 {
			m[nums[start]]--
			start++
		}
		res = max(res, end-start)
	}
	return res
}
