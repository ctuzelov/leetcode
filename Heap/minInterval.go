package Heap

import (
	"container/heap"
	"sort"
)

type Intervals struct {
	dur int
	end int
}

type minheapI []Intervals

func (h minheapI) Len() int { return len(h) }
func (h minheapI) Less(i, j int) bool {
	return h[i].dur < h[j].dur
}
func (h minheapI) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *minheapI) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Intervals))
}

func (h *minheapI) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func (h *minheapI) Top() Intervals {
	return (*h)[0]
}

func MinInterval(intervals [][]int, queries []int) []int {
	var res []int
	hp := new(minheapI)
	res = make([]int, len(queries))
	var qrs [][]int

	for i := 0; i < len(queries); i++ {
		qrs = append(qrs, []int{queries[i], i})
	}

	sort.SliceStable(qrs, func(i, j int) bool {
		{
			return qrs[i][0] < qrs[j][0]
		}
	})

	sort.SliceStable(intervals, func(i, j int) bool {
		{
			return intervals[i][0] < intervals[j][0]
		}
	})

	j := 0
	for _, n := range qrs {
		for j < len(intervals) && intervals[j][0] <= n[0] {
			heap.Push(hp, Intervals{intervals[j][1] - intervals[j][0] + 1, intervals[j][1]})
			j++
		}

		for hp.Len() > 0 && hp.Top().end < n[0] {
			heap.Pop(hp)
		}

		if hp.Len() > 0 {
			info := hp.Top()
			res[n[1]] = info.dur
		}
	}

	for i := 0; i < len(res); i++ {
		if res[i] == 0 {
			res[i] = -1
		}
	}

	return res

}

/*

[2,3],[2,5],[1,8],[20,25]  2,19,5,22

[2, 1], [5, 3], [19, 2], [22, 4]

[1,8,8], [2,3,2], [2,5,4], [20,25,6]

[2,3,2], [2,5,4], [20,25,6], [1,8,8]



*/
