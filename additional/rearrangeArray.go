package additional

func RearrangeArray(nums []int) []int {

	pos, neg := make([]int, len(nums)/2), make([]int, len(nums)/2)
	p, n := 0, 0
	for _, num := range nums {
		if num > 0 {
			pos[p] = num
			p++
		} else {
			neg[n] = num
			n++
		}
	}

	for i := 0; i < len(nums); i += 2 {
		nums[i] = pos[i/2]
		nums[i+1] = neg[i/2]
	}

	return nums
}
