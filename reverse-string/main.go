package main

// Дана строка input. Нужно написать функцию, которая перевернет строку справа налево.

import "fmt"

func main() {

	fmt.Println(reverseString("Привет я Петя!"))                  // !ятеП я тевирП
	fmt.Println(reverseString("Мой номер телефона: 89167854141")) // 14145879168 :анофелет ремон йоМ
	fmt.Println(reverseString("Hello, 世界"))                       // 界世 , olleH

}
