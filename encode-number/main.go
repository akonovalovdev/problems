package main

// Дано число `X` в десятичной системе счисления.
// Нужно написать функцию, которая переведет это число в двоичную систему счисления.
//
// А затем вывести на экран в формате `X (from) = __(2)`.

import "fmt"

func main() {
	fmt.Println(encodeNumber(123))  // 1111011
	fmt.Println(encodeNumber(100))  // 1100100
	fmt.Println(encodeNumber(0))    //
	fmt.Println(encodeNumber(1313)) // 10100100001
}
