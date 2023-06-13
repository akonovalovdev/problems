package main

import "fmt"

// Фраза является палиндромом, если после преобразования всех прописных букв в строчные
// и удаления всех не буквенно-цифровых символов она читается одинаково вперед и назад.
// Буквенно-цифровые символы включают буквы и цифры.
//
// Если задана строка s, возвращаем true если это палиндром, или false если иначе.
//

func main() {

	fmt.Println(isPalindrome("A man, a Plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car"))                     // false
	fmt.Println(isPalindrome(" "))                              // true
	fmt.Println(isPalindrome("0P"))
}
