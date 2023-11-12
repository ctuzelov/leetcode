package dp

import (
	"math"
	"sort"
)

func FindLongestChain(pairs [][]int) int {

	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})

	curr, mx := math.MinInt, 0
	for i := 0; i < len(pairs); i++ {
		if curr < pairs[i][0] {
			curr = pairs[i][1]
			mx++
		}
	}
	return mx
}
