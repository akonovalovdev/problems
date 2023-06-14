package main

import "fmt"

// Вам даны два массива целых чисел nums1 и nums2, отсортированные в неубывающем порядке,
// и два целых числа m и n, представляющие количество элементов в nums1 и nums2 соответственно.
//
// Объединить nums1 и nums2 в один массив, отсортированный в неубывающем порядке.
//
// Окончательный отсортированный массив не должен возвращаться функцией,
// а должен храниться внутри массива nums1 . Чтобы приспособиться к этому, nums1 имеет длину m + n,
// где первые m элементы обозначают элементы, которые должны быть объединены,
// а последние n элементы установлены 0 и должны игнорироваться. nums2 имеет длину n.

func main() {
	fmt.Println(merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)) // 1 2 2 3 5 6
	fmt.Println(merge([]int{1}, 1, []int{}, 0))                       // 1
	fmt.Println(merge([]int{0}, 0, []int{1}, 1))                      // 1
}
