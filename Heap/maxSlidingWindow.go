package Heap

import "container/heap"

type Pairs struct {
	index int
	num   int
}

type maxHeapP []Pairs

func (h maxHeapP) Len() int           { return len(h) }
func (h maxHeapP) Less(i, j int) bool { return h[i].num > h[j].num }
func (h maxHeapP) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeapP) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Pairs))
}

func (h *maxHeapP) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func (h *maxHeapP) Top() Pairs {
	return (*h)[0]
}

func MaxSlidingWindow(nums []int, k int) []int {
	var res []int

	hp := new(maxHeapP)
	l := 0

	for r := 0; r < len(nums); r++ {
		window := r - l + 1
		heap.Push(hp, Pairs{r, nums[r]})
		if window == k {
			for hp.Len() > 0 && hp.Top().index < l {
				heap.Pop(hp)
			}

			info := hp.Top()
			res = append(res, info.num)
			l++
		}
	}
	return res
}
