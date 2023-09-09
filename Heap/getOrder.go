package Heap

import (
	"container/heap"
	"sort"
)

type Task struct {
	proTime int
	index   int
}

type minheapT []Task

func (h minheapT) Len() int { return len(h) }
func (h minheapT) Less(i, j int) bool {
	return h[i].proTime < h[j].proTime || h[i].proTime == h[j].proTime && h[i].index < h[j].index
}
func (h minheapT) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *minheapT) Push(x interface{}) {
	*h = append(*h, x.(Task))
}

func (h *minheapT) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func (h *minheapT) Top() Task {
	return (*h)[0]
}

func GetOrder(tasks [][]int) []int {
	var res []int
	hp := new(minheapT)

	for i := 0; i < len(tasks); i++ {
		tasks[i] = append(tasks[i], i)
	}

	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i][0] == tasks[j][0] {
			return tasks[i][1] < tasks[j][1]
		}
		return tasks[i][0] < tasks[j][0]
	})

	sec, t_i := tasks[0][0], 0

	for t_i < len(tasks) || hp.Len() > 0 {
		// Add tasks to the heap if their arrival time matches the current time
		for t_i < len(tasks) && sec >= tasks[t_i][0] {
			heap.Push(hp, Task{tasks[t_i][1], tasks[t_i][2]})
			t_i++
		}

		// Check if the heap is empty
		if hp.Len() == 0 {
			sec = tasks[t_i][0]
		} else {
			info := heap.Pop(hp).(Task)
			res = append(res, info.index)
			sec += info.proTime
		}
	}

	return res
}
