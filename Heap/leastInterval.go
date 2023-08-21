package Heap

import (
	"container/heap"
	"strconv"
)

func LeastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	var (
		res      int
		entry    Counter
		taskFreq = map[string]int{}
		interval = map[string]int{}
		hp       = new(mxHeap)
	)

	for _, c := range tasks {
		taskFreq[string(c)]++
	}

	for k, v := range taskFreq {
		heap.Push(hp, Counter{char: k, freq: v})
	}

	for hp.Len() > 0 || len(interval) != 0 {
		if hp.Len() > 0 {
			entry = heap.Pop(hp).(Counter)
			interval[entry.char+strconv.Itoa(entry.freq-1)] = n + 1
		}

		for k := range interval {
			interval[k]--
			if interval[k] == 0 {
				freq, _ := strconv.Atoi(k[1:])
				if freq > 0 {
					heap.Push(hp, Counter{char: k[:1], freq: freq})
				}
				delete(interval, k)
			}
		}
		res++
	}

	return res - n
}

/*

	type Counter struct {
		char string
		freq int
	}

	type mxHeap []Counter

	func (h mxHeap) Len() int { return len(h) }

	func (h mxHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }

	func (h mxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

	func (h *mxHeap) Push(x interface{}) { (*h) = append((*h), x.(Counter)) }

	func (h *mxHeap) Pop() interface{} {
		toRemove := (*h)[len((*h))-1]
		(*h) = (*h)[:len((*h))-1]
		return toRemove
	}


*/
