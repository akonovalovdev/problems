package main

func constainsDuplicate(nums []int) bool {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		_, ok := m[nums[i]]
		if !ok {
			m[nums[i]] = 1
		} else {
			return true
		}
	}
	return false
}
