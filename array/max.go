package array

import (
	"sort"
)

func FindMaxBruteForce(arr []int) int {
	sort.Ints(arr)
	return arr[len(arr)-1]
}

func FindMaxOptimized(arr []int) int {
	max := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
