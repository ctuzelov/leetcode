package dp

import "math"

func MaxProduct(nums []int) int {
	mx, mn := 1, 1
	res := math.MinInt

	for _, n := range nums {
		if n == 0 {
			mn, mx = 1, 1
		}

		mx = Max(Max(n*mx, n*mn), n)
		mn = Min(Min(n*mx, n*mn), n)
	}

	return res
}

/*
2,3,-2,4,-3,-3,100,1000

mx = 2
mn =



*/
