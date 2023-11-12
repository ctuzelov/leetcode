package additional

func Square(arr []int) (res []int) {

	res = make([]int, len(arr))
	l, r := 0, len(arr)-1

	for i := len(arr) - 1; i >= 0 && l <= r; i-- {
		if arr[l] < 0 {
			if arr[l]*arr[l] > arr[r]*arr[r] {
				res[i] = arr[l] * arr[l]
				l++
				continue
			}
		}
		res[i] = arr[r] * arr[r]
		r--
	}

	return
}
