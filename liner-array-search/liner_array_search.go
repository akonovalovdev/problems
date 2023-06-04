package main

func linearSearch(nums []int, target int) int {
	const invalidIndex = -1
	for i, v := range nums {
		if v == target {
			return i
		}
	}
	return invalidIndex
}

// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == target {
// 			return i
// 		}
// 	}
// 	return invalidIndex
// }
