package array

func UniqueBruteForce(arr []int) {
	intMap := map[int]struct{}{}

	for i := 0; i < len(arr); i++ {
		intMap[arr[i]] = struct{}{}
	}

	idx := 0
	for k := range intMap {
		arr[idx] = k
		idx += 1
	}
}

func UniqueOptimized(arr []int) {
	secPtr := 0
	for i := 0; i < len(arr); i++ {
		if arr[secPtr] != arr[i] {
			secPtr += 1
			arr[secPtr] = arr[i]
		}
	}
}
