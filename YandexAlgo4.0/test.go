package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, element := range arr[1:] {
		if element <= pivot {
			left = append(left, element)
		} else {
			right = append(right, element)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	fmt.Println("Unsorted array:", arr)

	sortedArr := quickSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
