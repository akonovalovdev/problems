package main

// задание самая длинная последовательность

import (
	"fmt"
	"sort"
)

// Учитывая несортированный массив целых чисел nums,
// вернуть длину самой длинной последовательности последовательных элементов.
//
// Вы должны написать алгоритм, работающий во  O(n) времени.

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))         //4
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})) //9
	fmt.Println(longestConsecutive([]int{0, 6, 10, 2, 4, 8, 12}))        //-1
	fmt.Println(longestConsecutive([]int{}))                             //-1
}

func longestConsecutive(nums []int) int {
	const invalidCount = -1
	var maxCount, count, tracker int
	if len(nums) == 0 {
		return invalidCount
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		switch {
		case nums[i] > tracker:
			tracker = nums[i] + 1
			count = 1
		case nums[i] == tracker:
			tracker = nums[i] + 1
			count++
		}
		if count > maxCount {
			maxCount = count
		}
	}
	if maxCount > 1 {
		return maxCount
	}
	return invalidCount
}
