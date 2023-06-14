package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	// если один из массивов пустой, возвращаем сразу другой
	if m == 0 {
		nums1 = append(nums1[:m], nums2...)
		fmt.Println(nums1)
		return
	}
	if n == 0 {
		fmt.Println(nums1)
		return
	}
	// создаём переменные индексов для трёх массивов nums1, nums2 и nums 1 c общей длинной (n+m)
	i := m
	j := n - 1
	k := (m + n) - 1
	//итерируемся до тех пор пока не закончатся элементы во втором массиве(который вставляем)
	for j >= 0 {
		if i > 0 { // исключаем панику, если элементы в первом массиве все уже взяты (пример как в кейсе 4)
			if nums1[i-1] >= nums2[j] { // если элементы есть, тогда сверяем последний с последним элементов другого массива
				nums1[k] = nums1[i-1] // если первый больше или такой же, то вписываем в конец общей длины это значение
				i--                   // уменьшаем индекс последнего элемента из 1 массива
				k--                   // уменьшаем индекс последнего элемента общей длины массива
				continue              // завершаем итерацию
			}
		}
		// если первая проверка не прошла, значит значение последнего индекса второго массива больше первого или у первого
		// закончились элементы
		nums1[k] = nums2[j] // соответственно записываем в текущий конец общей длины данное значение из второго массива
		j--                 // уменьшаем индекс последнего элемента из 2 массива
		k--                 // уменьшаем индекс последнего элемента из общей длины массива
	}
	fmt.Println(nums1)
}

//	// исключаем кейс с пустым вторым срезом
//	if n == 0 {
//		fmt.Println(nums1)
//		return
//	}
//	//исключаем кейс с пустым первым срезом
//
//	if m == 0 {
//		nums1 = append(nums1[:m], nums2...)
//		fmt.Println(nums1)
//		return
//	}
//	mn := m + n
//	for _, v := range nums2 {
//		nums1 = insert(v, nums1, mn)
//	}
//	fmt.Println(nums1)
//}
//
//func insert(v int, nums1 []int, mn int) []int {
//	var firstValIndex = 0
//	var a []int             // создаём срез дубликат, чтобы корректно прибавлять куски до и после вставки нового элемента
//	a = append(a, nums1...) // добавляем все значения в слайс дубликат
//
//	for i := 0; i < mn; i++ {
//		// исключаем кейс, когда исходный массив начинается с отрицательного числа, чтобы пропустить
//		// впереди идущие нули
//
//		if firstValIndex >= nums1[i] && firstValIndex < v && len(nums1) > 1 {
//			continue
//		} else {
//			firstValIndex = -1000000
//		}
//
//		if nums1[i] >= v {
//			nums1 = append(nums1[:i], v)              // добавляем значение в соответсвующее место
//			nums1 = append(nums1[:i+1], a[i:mn-1]...) // дописываем всё что стояло после, без последнего нуля
//			return nums1
//		}
//		if nums1[i] == 0 {
//			nums1[i] = v
//			return nums1
//		}
//	}
//	return nums1
//}
