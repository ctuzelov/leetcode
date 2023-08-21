package Heap

import "container/heap"

func LongestDiverseString(a int, b int, c int) string {
	var (
		res      string
		prevChar string
		prevFreq int
		h        = &MaxHeap{}
	)

	heap.Push(h, charNum{"a", a})
	heap.Push(h, charNum{"b", b})
	heap.Push(h, charNum{"c", c})

	for h.Len() > 0 {
		entry := heap.Pop(h).(charNum)

		if entry.freq >= 2 && entry.freq > prevFreq {
			entry.freq -= 2
			res += entry.char + entry.char
		} else {
			entry.freq -= 1
			res += entry.char
		}

		if prevFreq > 0 {
			heap.Push(h, charNum{prevChar, prevFreq})
		}

		prevChar = entry.char
		prevFreq = entry.freq

	}
	return res
}

/*

	a = 7 b = 1 c = 1

	aabaacaa

	a = 7 b = 2 c = 1

	aabaabaaca

	a = 3 b = 3 c = 2

	aabbccab

*/
