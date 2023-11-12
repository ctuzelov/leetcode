package dp

func MaxSumTwoNoOverlap(nums []int, firstLen, secondLen int) int {
	prefixSum := make([]int, len(nums)+1)

	for i := 1; i < len(prefixSum); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}

	return max(GetSum(prefixSum, firstLen, secondLen), GetSum(prefixSum, secondLen, firstLen))
}

func GetSum(nums []int, firstLen, secondLen int) (res int) {

	firstMaxSum := 0

	for i := firstLen + secondLen; i < len(nums); i++ {
		firstSum := nums[i-secondLen] - nums[i-secondLen-firstLen]
		firstMaxSum = max(firstMaxSum, firstSum)
		secondSum := nums[i] - nums[i-secondLen]
		res = max(res, secondSum+firstMaxSum)
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
