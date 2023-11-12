package dp

func GoodDaysToRobBank(security []int, time int) []int {
	n := len(security)
	var ans []int

	if n < 2*time+1 {
		return ans
	}

	dp1 := make([]int, n)
	dp2 := make([]int, n)

	for i := 1; i < n-time; i++ {
		if security[i-1] >= security[i] {
			dp1[i] = dp1[i-1] + 1
		} else {
			dp1[i] = 0
		}
	}

	for j := n - 2; j >= time; j-- {
		if security[j+1] >= security[j] {
			dp2[j] = dp2[j+1] + 1
		} else {
			dp2[j] = 0
		}
	}

	for i := time; i < n-time; i++ {
		if dp1[i] >= time && dp2[i] >= time {
			ans = append(ans, i)
		}
	}

	return ans
}

/*
5,3,3,3,5,6,2


*/
