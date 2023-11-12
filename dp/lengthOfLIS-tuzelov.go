package dp

func LengthOfLIS(nums []int) int {
	lis := make([]int, len(nums))
	for i := range lis {
		lis[i] = 1
	}

	max := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				lis[i] = Max(lis[i], 1+lis[j])
				max = Max(max, lis[i])
			}
		}
	}
	return max
}
