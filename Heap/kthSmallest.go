package Heap

import (
	"container/heap"
	"strconv"
)

type Node struct {
	Sum     int
	Indexes []int
}

type minHeap []Node

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].Sum <= h[j].Sum
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *minHeap) Pop() interface{} {
	result := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return result
}

func getMarked(arr []int) string {
	res := strconv.Itoa(arr[0])
	for i := 1; i < len(arr); i++ {
		res += ":" + strconv.Itoa(arr[i])
	}
	return res
}

func getSum(mat [][]int, indexes []int) int {
	sum := 0

	for i, index := range indexes {
		sum += mat[i][index]
	}

	return sum
}

func KthSmallest(mat [][]int, k int) int {
	hp := new(minHeap)
	visited := make(map[string]bool)

	indexes := make([]int, len(mat))
	sum := getSum(mat, indexes)

	heap.Push(hp, Node{sum, indexes})

	for k > 0 {
		vals := heap.Pop(hp).(Node)

		path := getMarked(vals.Indexes)

		if visited[path] {
			continue
		}

		visited[path] = true
		sum = vals.Sum

		for i := 0; i < len(mat); i++ {
			if vals.Indexes[i]+1 >= len(mat[0]) {
				continue
			}

			indexes = append([]int{}, vals.Indexes...)
			indexes[i]++

			heap.Push(hp, Node{getSum(mat, indexes), indexes})
		}
		k--
	}
	return sum
}
