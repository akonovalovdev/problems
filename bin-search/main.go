package main

//Дан **отсортированный (упорядоченный от меньшего к большему)** массив целых чисел `nums` и целое число `target`.
//
//Необходимо проверить, есть ли в массиве заданное число, если есть - напечатать  его индекс, если нет, напечатать -1.

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
