package main

func constainsDuplicate(nums []int) bool {
	//создаём мапу, куда будем записывать и сравнивать уникальные элементы массива
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		_, ok := m[nums[i]] // проверяем есть ли в мапе уже такой ключ
		if !ok {
			m[nums[i]] = 1 // если ключ не найден, то записываем его
			continue
		}
		return true // если найден, то возвращаем правду

	}
	return false // если не найден за весь цикл, то возвращаем ложь
}
