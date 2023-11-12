package dp

func CountBits(n int) []int {
	res := make([]int, n+1)
	addit := 1

	for i := 1; i <= n; i++ {
		if addit*2 == i {
			addit *= 2
		}

		res[i] = res[i-addit] + 1
	}

	return res

}
