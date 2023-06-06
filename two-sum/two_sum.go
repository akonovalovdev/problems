package main

func twoSum(nums []int, target int) []int {
	// создайм мапу, где ключами будут разница искомой суммы и значения проверяемых индексов
	memory := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		raznitsa := target - nums[i]  // записываем разницу
		index, ok := memory[raznitsa] // проверяем существование ключа - разницы
		if ok {
			//если нашли то возвращаем текущий индекс и значение(индекс) того ключа сто нашли
			return []int{index, i}
		}
		// если не нашли то в значение мапы записываем текущий индекс
		memory[nums[i]] = i
	}
	return nil
}
