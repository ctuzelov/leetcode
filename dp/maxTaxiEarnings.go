package dp

import "sort"

func MaxTaxiEarnings(n int, rides [][]int) (res int64) {
	sort.SliceStable(rides, func(i, j int) bool {
		return rides[i][1] < rides[j][1]
	})
	return res
}

/*

[1,6,1],[3,10,2],[10,12,3],[11,12,2],[12,15,2],[13,18,1]
[1,6,1, 6],[3,10,2, 9],[10,12,3, 5],[11,12,2, 3],[12,15,2, 5],[13,18,1, 6]

*/