package dp

func MaxSubarraySumCircular(nums []int) int {
	n := len(nums)
	nonCircularSum := kanade(nums, n)
	total := 0
	for i := 0; i < n; i++ {
		total = total + nums[i]
		nums[i] = -nums[i]
	}
	circularSum := total + kanade(nums, n)
	if circularSum == 0 {
		return nonCircularSum
	}
	return Max(nonCircularSum, circularSum)
}

func kanade(arr []int, n int) int {
	max := -999999
	sum := 0
	for i := 0; i < n; i++ {
		sum = Max(arr[i], arr[i]+sum)
		max = Max(max, sum)
	}
	return max
}

// func Max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

//[5,-3,5] 7
// [-5,3,-5] 3
