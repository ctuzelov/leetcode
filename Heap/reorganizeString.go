package Heap

import "container/heap"

type charNum struct {
	char string
	freq int
}

type MaxHeap []charNum

func (h MaxHeap) Less(i, j int) bool {
	return h[i].freq > h[j].freq
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h *MaxHeap) Pop() interface{} {
	l := len(*h)
	toRemove := (*h)[l-1]
	*h = (*h)[0 : l-1]
	return toRemove
}

func (h *MaxHeap) Push(x interface{}) {
	(*h) = append((*h), x.(charNum))
}

func ReorganizeString(s string) string {
	var (
		res      string
		m        = map[string]int{}
		prevChar string
		prevFreq int
		h        = &MaxHeap{}
	)

	for _, c := range s {
		m[string(c)]++
	}

	for k, v := range m {
		heap.Push(h, charNum{k, v})
	}

	for h.Len() > 0 {
		val := heap.Pop(h).(charNum)
		if prevFreq > 0 {
			heap.Push(h, charNum{prevChar, prevFreq})
		}
		prevChar = val.char
		val.freq -= 1
		res += prevChar
		prevFreq = val.freq
	}
	if len(res) != len(s) {
		return ""
	}
	return res
}
