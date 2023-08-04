package Heap

import "container/heap"

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func LastStoneWeight(stones []int) int {
	mxHeap := maxHeap(stones)
	heap.Init(&mxHeap)

	for len(mxHeap) >= 2 {
		x1 := heap.Pop(&mxHeap).(int)
		x2 := mxHeap[0]

		if x1 == x2 {
			_ = heap.Pop(&mxHeap).(int)
		} else {
			mxHeap[0] = x1 - x2
			heap.Fix(&mxHeap, 0)
		}
	}

	if len(mxHeap) == 0 {
		return 0
	}

	return heap.Pop(&mxHeap).(int)
}
