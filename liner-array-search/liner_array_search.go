package main

func linearSearch(nums []int, target int) int {
	const invalidIndex = -1  // выносим дефолтное значение в константу
	for i, v := range nums { // пробегаемся по массиву рэнджом
		if v == target { // сравниваем значение с таргетом
			return i // возвращаем индекс
		}
	}
	return invalidIndex // возвращаем дефолтное значение
}

// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == target {
// 			return i
// 		}
// 	}
// 	return invalidIndex
// }
