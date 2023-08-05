package Heap

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
*/

func Newt() *IntHeap {
	var items = make(IntHeap, 0)
	return &items
}

func MergeKLists(lists []*ListNode) *ListNode {

	var (
		res   = new(ListNode)
		mHeap = Newt()
	)
	heap.Init(mHeap)

	for _, v := range lists {
		for v != nil {
			heap.Push(mHeap, v.Val)
			v = v.Next
		}
	}

	mHeap.convert(res)
	return res.Next
}

func insert(root *ListNode, num int) *ListNode {
	temp := &ListNode{num, nil}
	if root == nil {
		root = temp
		return root
	}
	curr := root
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = temp
	return root
}

func (arr *IntHeap) convert(root *ListNode) {
	for len(*arr) > 0 {
		root = insert(root, heap.Pop(arr).(int))
	}
}
