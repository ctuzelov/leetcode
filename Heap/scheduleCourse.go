package Heap

import (
	"container/heap"
	"sort"
)

func ScheduleCourse(courses [][]int) int {
	hp := new(IntHeap)
	var cum int
	sort.SliceStable(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	for i := 0; i < len(courses); i++ {
		dur, end := courses[i][0], courses[i][1]
		cum += dur
		heap.Push(hp, dur)

		if cum > end {
			cum -= heap.Pop(hp).(int)
		}
	}

	return hp.Len()
}

/*
[100,200],[200,1300],[1000,1250],[2000,3200]
*/
