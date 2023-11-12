package dp

func Rob(nums []int) int {
	var r1, r2 int

	for i := 0; i < len(nums); i++ {
		temp := Max(nums[i]+r1, r2)
		r1 = r2
		r2 = temp
	}
	return r2
}

/*
2, 7, 9, 3, 1
*/

// func Max(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }
