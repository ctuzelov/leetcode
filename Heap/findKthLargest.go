package Heap

import "container/heap"

/*
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
*/

func FindKthLargest(nums []int, k int) int {
	var (
		res    int
		mxHeap = IntHeap(nums)
	)

	heap.Init(&mxHeap)

	for k > 0 {
		res = heap.Pop(&mxHeap).(int)
		k--
	}
	return res
}
