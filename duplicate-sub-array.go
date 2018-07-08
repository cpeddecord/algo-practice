package main

// DuplicateArr passes a window over an array of integers and finds duplicates
// within each sub array
//
//   examples:
// arr{1, 2, 3, 4, 1, 5}, k=4
// no duplicates in consecutive subarrays:
//   [1, 2, 3, 4] [2, 3, 4, 1] [3, 4, 1, 5]
//
// arr{1, 2, 3, 2, 1}, k=3
// duplicate in consecutive subarray:
//   [1, 2, 3] [2, 3, 2] [3, 2, 1]
func DuplicateArr(arr []int, k int) bool {
	// cache our moving window
	nums := make(map[int]int)

	for i, n := range arr {
		// ensure our window stays current
		if i >= k {
			nums[arr[i-k]] = 0
		}

		if nums[n] == n {
			return true
		}

		nums[n] = n
	}

	return false
}
