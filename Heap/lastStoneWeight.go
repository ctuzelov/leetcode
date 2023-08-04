package Heap

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func LastStoneWeight(stones []int) int {
	mxHeap := IntHeap(stones)
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
