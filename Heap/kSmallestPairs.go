package Heap

import (
	"container/heap"
	"strconv"
)

func KSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var res [][]int
	minheap := new(MinHeap)
	m := make(map[string]bool)
	heap.Push(minheap, []int{0, 0, nums1[0], nums2[0]})
	for k > 0 && minheap.Len() > 0 {
		val := heap.Pop(minheap).([]int)

		l, r := val[0], val[1]
		res = append(res, []int{val[2], val[3]})

		visitedL := strconv.Itoa(l+1) + "and" + strconv.Itoa(r)
		visitedR := strconv.Itoa(l) + "and" + strconv.Itoa(r+1)

		if l+1 < len(nums1) && !m[visitedL] {
			m[visitedL] = true
			heap.Push(minheap, []int{l + 1, r, nums1[l+1], nums2[r]})
		}

		if r+1 < len(nums2) && !m[visitedR] {
			m[visitedR] = true
			heap.Push(minheap, []int{l, r + 1, nums1[l], nums2[r+1]})
		}
		k--
	}
	return res
}

type MinHeap [][]int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i][2]+h[i][3] < h[j][2]+h[j][3]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	toRemove := (*h)[len((*h))-1]
	(*h) = (*h)[:len((*h))-1]
	return toRemove
}
