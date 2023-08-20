package Heap

import "container/heap"

type Counter struct {
	num  int
	char string
	freq int
}

type mxHeap []Counter

func (h mxHeap) Len() int {
	return len(h)
}

func (h mxHeap) Less(i, j int) bool {
	return h[i].freq > h[j].freq
}

func (h mxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *mxHeap) Push(x interface{}) {
	(*h) = append((*h), x.(Counter))
}

func (h *mxHeap) Pop() interface{} {
	toRemove := (*h)[len((*h))-1]
	(*h) = (*h)[:len((*h))-1]
	return toRemove
}

func RearrangeBarcodes(b []int) []int {
	var res []int
	m := make(map[int]int)
	hp := &mxHeap{}
	prevNum := 0
	prevFreq := 0

	for _, num := range b {
		m[num]++
	}

	for k, v := range m {
		heap.Push(hp, Counter{num: k, freq: v})
	}

	for hp.Len() > 0 {
		entry := heap.Pop(hp).(Counter)

		if prevFreq > 0 {
			heap.Push(hp, Counter{num: prevNum, freq: prevFreq})
		}

		res = append(res, entry.num)
		prevNum = entry.num
		prevFreq = entry.freq
		prevFreq--
	}
	return res
}
