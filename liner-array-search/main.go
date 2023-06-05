package main

// Дан массив целых чисел `nums` и целое число `target`.
//
// Необходимо проверить, есть ли в массиве заданное число, если
// есть - напечатать  его индекс, если нет, напечатать -1.

import (
	"fmt"
)

func main() {
	fmt.Println(linearSearch([]int{5, 7, 8}, 7))                                   // 1
	fmt.Println(linearSearch([]int{5, 7, 8}, 3))                                   // -1
	fmt.Println(linearSearch([]int{6, 6, 7}, 1))                                   // -1
	fmt.Println(linearSearch([]int{5, 7, 8, 12, 43, 54, 325, 765, 2, 6546, 3}, 3)) // 10
	fmt.Println(linearSearch([]int{6, 6, 7, 9, 0, 2, 3, 46, 77, 88, 3}, 1))        // -1
}
