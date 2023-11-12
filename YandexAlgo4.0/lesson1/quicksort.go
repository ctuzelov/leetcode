package lesson1

func QuickSort(arr []int) []int {
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

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, pivot), right...)
}
