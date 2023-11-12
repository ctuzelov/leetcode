package dp

import (
	"container/heap"
)

type IntHeap [][2]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Top() interface{}   { return h[0] }

func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func MaxScoreSightseeingPair(values []int) int {
	var res int
	hp := new(IntHeap)

	for i := 1; i < len(values); i++ {
		heap.Push(hp, [2]int{i, values[i] - i})
	}

	for i := range values {
		n := values[i]

		for hp.Len() > 0 && i >= hp.Top().([2]int)[0] {
			heap.Pop(hp)
		}

		if hp.Len() > 0 {
			res = max(res, i+n+hp.Top().([2]int)[1])
		}
	}

	return res
}
