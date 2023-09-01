package Heap

import (
	"container/heap"
	"math"
)

type intHeap [][2]int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([2]int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *intHeap) Top() interface{} {
	return (*h)[0]
}

func SmallestRange(nums [][]int) []int {
	h := &intHeap{}
	heap.Init(h)
	maxVal := math.MinInt32
	for i, list := range nums {
		heap.Push(h, [2]int{i, list[0]})
		if list[0] > maxVal {
			maxVal = list[0]
		}
	}
	covered := make([]int, len(nums))
	for i := range covered {
		covered[i] = 0
	}
	res := []int{(*h)[0][1], maxVal}
	for {
		top := heap.Pop(h).([2]int)
		covered[top[0]]++
		if covered[top[0]] == len(nums[top[0]]) {
			break
		}
		if nums[top[0]][covered[top[0]]] > maxVal {
			maxVal = nums[top[0]][covered[top[0]]]
		}
		heap.Push(h, [2]int{top[0], nums[top[0]][covered[top[0]]]})
		if maxVal-(*h)[0][1] < res[1]-res[0] {
			res = []int{(*h)[0][1], maxVal}
		}
	}
	return res
}

/*
[[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]
[0,0,0]
[0,1,0]

*/
