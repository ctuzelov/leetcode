package Heap

import "container/heap"

type Available struct {
	weight int
	index  int
}

type Unavailable struct {
	time   int
	weight int
	index  int
}

type minheapA []Available
type minheapU []Unavailable

func (h minheapA) Len() int { return len(h) }
func (h minheapA) Less(i, j int) bool {
	if h[i].weight == h[j].weight {
		return h[i].index < h[j].index
	} else {
		return h[i].weight < h[j].weight
	}
}
func (h minheapA) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *minheapA) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Available))
}

func (h *minheapA) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func (h *minheapA) Top() Available {
	return (*h)[0]
}

func (h minheapU) Len() int           { return len(h) }
func (h minheapU) Less(i, j int) bool { return h[i].time < h[j].time }
func (h minheapU) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minheapU) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Unavailable))
}

func (h *minheapU) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func (h *minheapU) Top() Unavailable {
	return (*h)[0]
}

func AssignTasks(servers []int, tasks []int) []int {
	var (
		res  []int
		avlb = new(minheapA)
		unlb = new(minheapU)
		t_i  int
	)

	for i, w := range servers {
		heap.Push(avlb, Available{w, i})
	}

	for i := 0; ; i++ {
		if len(res) == len(tasks) {
			return res
		}

		if len(servers) == unlb.Len() && unlb.Top().time-i != 0 {
			continue
		}

		for unlb.Len() > 0 && unlb.Top().time-i == 0 {
			info := heap.Pop(unlb).(Unavailable)
			heap.Push(avlb, Available{info.weight, info.index})
		}

		info := heap.Pop(avlb).(Available)

		res = append(res, info.index)
		heap.Push(unlb, Unavailable{i + tasks[t_i], info.weight, info.index})
		t_i++
	}

}

/*

[26,50,47,11,56,31,18,55,32,9,4,2,23,53,43,0,44,30,6,51,29,51,15,17,22,34,38,33,42,3,25,10,49,51,7,58,16,21,19,31,19,12,41,35,
45,52,13,59,47,36,1,28,48,39,24,8,46,20,5,54,27,37,14,57,40,59,8,45,4,51,47,7,58,4,31,23,54,7,9,56,2,46,56,1,17,42,11,30,12,
44,14,32,7,10,23,1,29,27,6,10,33,24,19,10,35,30,35,10,17,49,50,36,29,1,48,44,7,11,24,57,42,30,10,55,3,20,38,15,7,46,32,21,40,
16,59,30,53,17,18,22,51,11,53,36,57,26,5,56,36,55,31,34,57,7,52,37,31,10,0,51,41,2,32,25,0,7,49,47,13,14,24,57,28,4,45,43,39,
38,8,2,44,45,29,25,25,12,54,5,44,30,27,23,26,7,33,58,41,25,52,40,58,9,52,40]

[26,50,47,11,56,31,18,55,32,9,4,2,23,53,43,0,44,30,6,51,29,51,15,17,22,34,38,33,42,3,25,10,49,51,7,58,16,21,19,31,19,12,41,35,
45,52,13,59,47,36,1,28,48,39,24,8,46,20,5,54,27,37,14,57,40,59,8,45,4,51,47,7,58,4,31,23,54,7,9,56,2,46,56,1,17,42,11,30,12,
44,14,32,7,10,23,1,29,27,6,10,33,24,19,10,35,30,35,10,17,49,50,36,29,1,48,44,7,11,24,57,42,30,10,55,3,20,38,15,7,46,32,21,40,
16,59,30,53,17,18,22,51,11,53,36,57,26,5,36,56,55,31,34,57,7,52,37,31,10,0,51,41,2,32,25,0,7,49,47,13,14,24,57,28,4,45,43,39,
38,8,2,44,45,29,25,25,12,54,5,44,30,27,23,26,7,33,58,41,25,52,40,58,9,52,40]

*/
