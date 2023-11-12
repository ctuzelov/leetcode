package backtracking

import "fmt"

func MaximumSumOfHeights(maxHeights []int) int64 {
	// Define a function to find the peaks and calculate maximum sum of heights
	findPeak := func(arr []int) []int {
		stack := make([][2]int, 0)
		acc := 0
		peeks := make([]int, 0)

		for _, height := range arr {
			count := 1
			for len(stack) > 0 && height <= stack[len(stack)-1][0] {
				h, c := stack[len(stack)-1][0], stack[len(stack)-1][1]
				stack = stack[:len(stack)-1]
				acc -= h * c
				count += c
			}
			stack = append(stack, [2]int{height, count})
			acc += height * count
			peeks = append(peeks, acc)
		}
		return peeks
	}

	// Precompute sums from both tails (forward and reverse)
	fwd := make([]int, 1)
	fwd = append(fwd, findPeak(maxHeights)...)

	revHeights := make([]int, len(maxHeights))
	copy(revHeights, maxHeights)
	reverseSlice(revHeights)
	rev := findPeak(revHeights)
	reverseSlice(rev)
	rev = append(rev, 0)

	// Find the maximum sum for two tails
	var maxSum int64
	for i := range fwd {
		if int64(fwd[i]+rev[i]) > maxSum {
			maxSum = int64(fwd[i] + rev[i])
		}
	}

	fmt.Println(maxHeights, fwd)

	return maxSum
}

func reverseSlice(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
