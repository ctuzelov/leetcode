package dp

func WiggleMaxLength(nums []int) int {
	diff := []int{}
	for i := 0; i < len(nums)-1; i++ {
		diff = append(diff, nums[i+1]-nums[i])
	}
	ans := 0
	if len(diff) <= 1 {
		if len(diff) == 0 {
			return 1
		}
		if diff[0] == 0{
			return 1
		}
		return len(nums)
	}
	cur := diff[0]
	ans = 1
	if diff[0] == 0 {
		ans = 0
	}
	for i := 1; i < len(diff); i++ {
		if cur == 0 && diff[i] != 0 {
			cur = diff[i]
			ans = 1
			continue
		}
		if diff[i]*cur >= 0 {
			continue
		}
		ans += 1
		cur = diff[i]
	}
	return ans + 1
}
