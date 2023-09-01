package Heap

import (
	"container/heap"
	"sort"
)

func MinRefuelStops(target int, startFuel int, stations [][]int) int {
	hp := new(IntHeap)

	sort.SliceStable(stations, func(i, j int) bool {
		return stations[i][0] < stations[j][0]
	})

	res := 0

	for i := 0; i <= len(stations); i++{
		
		dist := 0

		if i < len(stations){
			dist = stations[i][0]
		}

		for startFuel < dist{
			if hp.Len() == 0{
				return -1
			}

			startFuel += heap.Pop(hp).(int)
			res++
		}
		
		if i < len(stations){
			heap.Push(hp, stations[i][1])
		}
	}
	return res
}

/*
target = 100, startFuel = 10, stations = [[10,60],[20,30],[30,30],[60,40]]
1. f 60 p 10
2. f 70 p 30


.....[10,60].....[20,30].....[30,30].........[60,40]

    ->



*/
