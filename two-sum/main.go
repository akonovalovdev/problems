package main

// Дан массив целых чисел `nums` и целое число `target` .
//
// Нужно найти индексы двух чисел из данного массива, сумма которых будет равна `target`.
//
// Важные условия задачи:
//
// - Каждый данный пример входных данных будет иметь ровно одно решен
// - Нельзя использовать один и тот же элемент массива дважды
// - Можно выдать ответ в любом порядке и форме.

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{1, 2, 3}, 5))                           // [1,2]
	fmt.Println(twoSum([]int{3, 3}, 6))                              // [0,1]
	fmt.Println(twoSum([]int{9, 5, 4, 3, 8, 7}, 9))                  // [1,2]
	fmt.Println(twoSum([]int{9, 5, 1, 4, 3, 8, 7, 3, 0, 5, 1}, 2))   // [2,10]
	fmt.Println(twoSum([]int{91, 53, 42, 32, 82, 7, 8, 45, 95}, 99)) // [0,6]
}
