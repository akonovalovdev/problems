package main

import "fmt"

// Вам даны два массива целых чисел nums1 и nums2, отсортированные в неубывающем порядке,
// и два целых числа m и n, представляющие количество элементов в nums1 и nums2 соответственно.
//
// Объединить nums1 и nums2 в один массив, отсортированный в неубывающем порядке.
//
// Окончательный отсортированный массив не должен возвращаться функцией,
// а должен храниться внутри массива nums1. Чтобы приспособиться к этому, nums1 имеет длину m + n,
// где первые m элементы обозначают элементы, которые должны быть объединены,
// а последние n элементы установлены 0 и должны игнорироваться. nums2 имеет длину n.

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)                  // 1 2 2 3 5 6
	merge([]int{1}, 1, []int{}, 0)                                        // 1
	merge([]int{0}, 0, []int{1}, 1)                                       // 1
	merge([]int{2, 0}, 1, []int{1}, 1)                                    // 1 2
	merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)                  // 1 2 3 4 5 6
	merge([]int{-1, 0, 0, 3, 3, 3, 0, 0, 0}, 6, []int{1, 2, 2}, 3)        // -1 0 0 1 2 2 3 3 3
	merge([]int{-1, 0, 0, 3, 3, 3, 0, 0, 0}, 6, []int{-1, 1, 2}, 3)       // -1 -1 0 0 1 2 3 3 3
	merge([]int{0, 0, 0, 0, 0}, 0, []int{1, 2, 3, 4, 5}, 5)               // 1 2 3 4 5
	merge([]int{-1, 1, 0, 0, 0, 0, 0, 0}, 2, []int{-1, 0, 1, 1, 2, 3}, 6) // -1 -1 0 1 1 1 2 3
	merge([]int{-1, 0, 0, 3, 3, 3, 0, 0, 0, 0}, 6, []int{-1, 0, 1, 4}, 4) // -1 -1 0 0 0 1 2 3 3 3
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		nums1 = append(nums1[:m], nums2...)
		fmt.Println(nums1)
		return
	}
	if n == 0 {
		fmt.Println(nums1)
		return
	}
	i := m
	j := n - 1
	k := (m + n) - 1
	for j >= 0 {
		if i > 0 {
			if nums1[i-1] >= nums2[j] {
				nums1[k] = nums1[i-1]
				i--
				k--
				continue
			}
		}
		nums1[k] = nums2[j]
		if j == 0 {
			break
		}
		j--
		k--
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
