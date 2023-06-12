package backtracking

import (
	"sort"
)

func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var backtrack func(ind int, pathsum int, path []int)

	backtrack = func(ind int, pathsum int, path []int) {
		if pathsum == target {
			cp := make([]int, len(path))
			copy(cp, path)
			res = append(res, cp)
			return
		}

		if pathsum > target {
			return
		}

		for i := ind; i < len(candidates); i++ {
			path = append(path, candidates[i])
			backtrack(i, pathsum+candidates[i], path)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, 0, []int{})
	return res
}
