package Heap

import "container/heap"

func FurthestBuilding(heights []int, bricks int, ladders int) int {
	h := new(IntHeap)
	heap.Init(h)

	res := len(heights)
	if res == 1 {
		return 0
	}

	for i := 1; i < len(heights); i++ {
		d := heights[i] - heights[i-1]

		if d > 0 {
			if bricks >= d {
				bricks -= d
				heap.Push(h, d)
			} else {
				if ladders == 0 {
					return i - 1
				}

				ladders--
				heap.Push(h, d)
				temp := heap.Pop(h).(int)

				if temp != d {
					bricks += temp - d
				}
			}
		}
	}
	return res
}

/*

	4,12,2,7,3,18,20,3,19  10 2

*/
