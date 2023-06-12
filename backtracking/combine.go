package backtracking

func Combine(n int, k int) [][]int {
	var backtrack func(ind int, temp []int)
	var res [][]int
	backtrack = func(ind int, temp []int) {
		if len(temp) == k {
			cp := make([]int, len(temp))
			copy(cp, temp)
			res = append(res, cp)
			return
		}
		for i := ind; i <= n; i++ {
			temp = append(temp, i)
			backtrack(i+1, temp)
			temp = temp[:len(temp)-1]
		}
	}
	backtrack(1, []int{})
	return res
}
