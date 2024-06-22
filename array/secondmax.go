package array

import "sort"

func FindSecondMaxBruteForce(arr []int) int {
	sort.Ints(arr)
	max := arr[len(arr)-1]
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != max {
			return arr[i]
		}
	}
	return -1
}
func FindSecondMaxBetter(arr []int) int {
	max := FindMaxOptimized(arr)
	secondMax := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] > secondMax && arr[i] != max {
			secondMax = arr[i]
		}
	}
	return secondMax
}

func FindSecondMaxOptimized(arr []int) int {
	max := arr[0]
	secondMax := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			secondMax = max
			max = arr[i]
		}
		if arr[i] > secondMax && arr[i] < max {
			secondMax = arr[i]
		}
	}

	return secondMax
}
