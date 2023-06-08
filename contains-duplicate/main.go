package main

import "fmt"

// Для заданного целочисленного массива nums возвращайте значение,
// true если какое-либо значение встречается в массиве не менее двух раз
//  значение false возвращайте если все элементы различны.

func main() {
	fmt.Println(constainsDuplicate([]int{1, 2, 3, 1}))                   // true
	fmt.Println(constainsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})) // true
	fmt.Println(constainsDuplicate([]int{1, 2, 3, 4}))                   // false

}
