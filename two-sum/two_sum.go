package main

func twoSum(nums []int, target int) []int {
	memory := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		raznitsa := target - nums[i]
		index, ok := memory[raznitsa]
		if ok {
			return []int{index, i}
		}
		memory[nums[i]] = i
	}
	return nil
}
