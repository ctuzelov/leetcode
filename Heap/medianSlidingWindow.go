package Heap

func MedianSlidingWindow(nums []int, k int) []float64 {
	var res []float64

	var minHeap minH
	var maxHeap maxH

	for i, num := range nums {
		if len(maxHeap) == 0 || num <= maxHeap[0] {
			maxHeap.add(num)
		} else {
			minHeap.add(num)
		}

		if len(maxHeap) > len(minHeap)+1 {
			minHeap.add(maxHeap.pop())
		} else if len(minHeap) > len(maxHeap) {
			maxHeap.add(minHeap.pop())
		}

		if i-k+1 >= 0 {
			removedIndex := i - k

			if removedIndex >= 0 {
				if nums[removedIndex] <= maxHeap[0] {
					maxHeap.remove(nums[removedIndex])
				} else {
					minHeap.remove(nums[removedIndex])
				}

				if len(maxHeap) > len(minHeap)+1 {
					minHeap.add(maxHeap.pop())
				} else if len(minHeap) > len(maxHeap) {
					maxHeap.add(minHeap.pop())
				}
			}

			var median float64

			if len(minHeap) == len(maxHeap) {
				median = float64(minHeap[0]+maxHeap[0]) / float64(2)
			} else {
				median = float64(maxHeap[0])
			}

			res = append(res, median)
		}
	}

	return res
}

type minH []int

type maxH []int

func (h *maxH) remove(item int) {
	for i := 0; i < len(*h); i++ {
		if item == (*h)[i] {
			(*h)[i] = (*h)[len(*h)-1]
			*h = (*h)[:len(*h)-1]

			if i == 0 {
				h.heapUp(0)
				return
			}

			parent := (i - 1) / 2

			if parent >= 0 && i < len(*h) {
				h.heapDown(i)
				h.heapUp(i)
				return
			}
		}
	}
}

func (h *minH) remove(item int) {
	for i := 0; i < len(*h); i++ {
		if item == (*h)[i] {
			(*h)[i] = (*h)[len(*h)-1]
			*h = (*h)[:len(*h)-1]

			if i == 0 {
				h.heapDown(0)
				return
			}

			parent := (i - 1) / 2

			if parent >= 0 && i < len(*h) {
				h.heapDown(i)
				h.heapUp(i)
				return
			}
		}
	}
}

func (h *maxH) pop() int {
	poppedItem := (*h)[0]

	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.heapDown(0)

	return poppedItem
}

func (h *minH) pop() int {
	poppedItem := (*h)[0]

	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.heapDown(0)

	return poppedItem
}

func (h *minH) add(num int) {
	*h = append(*h, num)
	h.heapUp(len(*h) - 1)
}

func (h *maxH) add(num int) {
	*h = append(*h, num)
	h.heapUp(len(*h) - 1)
}

func (h *minH) heapDown(p int) {
	l, r := 2*p+1, 2*p+2
	smaller := p

	if l < len(*h) && (*h)[l] < (*h)[smaller] {
		smaller = l
	}

	if r < len(*h) && (*h)[r] < (*h)[smaller] {
		smaller = r
	}

	if smaller != p {
		(*h)[smaller], (*h)[p] = (*h)[p], (*h)[smaller]
		h.heapDown(smaller)
	}
}

func (h *minH) heapUp(p int) {
	parent := (p - 1) / 2

	if parent >= 0 && (*h)[p] < (*h)[parent] {
		(*h)[p], (*h)[parent] = (*h)[parent], (*h)[p]
		h.heapUp(parent)
	}
}

func (h *maxH) heapDown(p int) {
	l, r := 2*p+1, 2*p+2
	bigger := p

	if l < len(*h) && (*h)[l] > (*h)[bigger] {
		bigger = l
	}

	if r < len(*h) && (*h)[r] > (*h)[bigger] {
		bigger = r
	}

	if bigger != p {
		(*h)[bigger], (*h)[p] = (*h)[p], (*h)[bigger]
		h.heapDown(bigger)
	}
}

func (h *maxH) heapUp(p int) {
	parent := (p - 1) / 2

	if parent >= 0 && (*h)[p] > (*h)[parent] {
		(*h)[p], (*h)[parent] = (*h)[parent], (*h)[p]
		h.heapUp(parent)
	}
}
