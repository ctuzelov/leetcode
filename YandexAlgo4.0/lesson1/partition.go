package lesson1

func Partition(arr []int, x int) (less int, more int) {

	for _, n := range arr {
		if n < x {
			less++
		} else {
			more++
		}
	}

	return
}
