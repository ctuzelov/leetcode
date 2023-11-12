package additional

import (
	"math/bits"
	"sort"
)

type num_bit struct {
	num int
	bit int
}

func SortByBits(arr []int) []int {
	data := make([]num_bit, len(arr), len(arr))

	for pos, v := range arr {
		data[pos] = num_bit{v, bits.OnesCount16(uint16(v))}
	}

	sort.SliceStable(data, func(i, j int) bool {
		return data[i].bit < data[j].bit
	})

	out := make([]int, len(arr), len(arr))
	for pos, v := range data {
		out[pos] = v.num
	}
	return out
}
