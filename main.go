package main

import (
	"fmt"

	"github.com/Tanmay451/dsa/array"
)

func main() {
	arr := []int{1, 2, 3, 4}

	fmt.Println(array.FindMaxBruteForce(arr))
	fmt.Println(array.FindMaxOptimized(arr))
	fmt.Println(array.FindSecondMaxBruteForce(arr))
	fmt.Println(array.FindSecondMaxBetter(arr))
	fmt.Println(array.FindSecondMaxOptimized(arr))

	sortedArry := []int{1, 1, 2, 2, 4, 4, 4}

	testCopy := make([]int, len(sortedArry))
	copy(testCopy, sortedArry)
	array.UniqueBruteForce(testCopy)
	fmt.Println(testCopy)

	copy(testCopy, sortedArry)
	array.UniqueOptimized(testCopy)
	fmt.Println(testCopy)

}
