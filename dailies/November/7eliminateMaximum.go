package November

import (
	"math"
	"sort"
)

func eliminateMaximum(dist []int, speed []int) int {
	times := make([]int, len(dist))

	for i := range times {
		times[i] = int(math.Ceil(float64(dist[i]) / float64(speed[i])))
	}

	sort.Ints(times)

	idx := 0
	for _, item := range times {
		if item <= idx {
			break
		}
		idx++
	}

	return idx
}
