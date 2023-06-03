package main

import "fmt"

func main() {
	fmt.Println(binSearch([]int{1, 2, 3, 4, 5, 7, 9, 10, 11}, 10)) // 7
	fmt.Println(binSearch([]int{1, 2, 3, 4, 5, 7, 9}, 5))          // 4
	fmt.Println(binSearch([]int{3, 4, 5, 7, 9}, 9))                // 4
	fmt.Println(binSearch([]int{5, 7, 9, 11}, 11))                 // 3
	fmt.Println(binSearch([]int{1, 2, 3, 4, 5, 7, 9, 10, 11}, 0))  // -1
	fmt.Println(binSearch([]int{1, 2, 3, 4, 5, 7, 9}, -100))       // -1
	fmt.Println(binSearch([]int{3, 4, 5, 7, 9}, 100))              // -1
	fmt.Println(binSearch([]int{5, 7, 9, 11}, 8))                  // -1
}

// функция выполняет поиск в слайсе заданного числа за логорифмическое время
func binSearch(arr []int, target int) int {
	const invalidIndex = -1 // дефолтный кейс
	l, r := 0, len(arr)-1   //l,r (left,right) -переменные "ЛЕВО" и "ПРАВО", обозначающие границы проверяемого отрезка
	for l < r {             // пока лево меньше права, ещё есть отрезок где необходимо осуществить проверку соответствия
		center := (l + r) / 2      // для поиска центра делим расстояние на 2
		if arr[center] == target { //проверка центра, если равен таргету, тогда сразу возвращаем
			return center
		}
		if arr[center] > target && center != l { //если центр больше таргета, и не равен левому краю
			r = center - 1 // правую границу сдвигаем до центра и отступаем на -1, так как он уже проверен
		} else if center != r { // проверяя левый край, но не указываем проверку центр меньше, тк оно одно оставшееся
			l = center + 1
		}
	}
	if arr[l] == target { // после завершения цикла важно перепроверить последний индекс, тк при делении они могут стать равными
		return l // так как у l и r одинаковые индексы, достаточно проверки одного
	}
	return invalidIndex
}
