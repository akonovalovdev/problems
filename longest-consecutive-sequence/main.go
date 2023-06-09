package main

import "fmt"

// Учитывая несортированный массив целых чисел nums,
// вернуть длину самой длинной последовательности последовательных элементов.
//
// Вы должны написать алгоритм, работающий во  O(n) времени.

longest-consecutive-sequence

func main() {
	fmt.Println(longestConsecutive([]int {100,4,200,1,3,2})) //4
	fmt.Println(longestConsecutive([]int {0,3,7,2,5,8,4,6,0,1})) //9
	fmt.Println(longestConsecutive([]int {0,3,7,2,5,8,4,6,0,1})) //-1
	fmt.Println(longestConsecutive([]int {})) //-1
}


