package dp

func Rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	return Max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

func rob(nums []int) int {
	var r1, r2 int

	for i := 0; i < len(nums); i++ {
		temp := Max(nums[i]+r1, r2)
		r1 = r2
		r2 = temp
	}
	return r2
}
