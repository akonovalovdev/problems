package main

import "fmt"

// Фраза является палиндромом, если после преобразования всех прописных букв в строчные
// и удаления всех не буквенно-цифровых символов она читается одинаково вперед и назад.
// Буквенно-цифровые символы включают буквы и цифры.
//
// Если задана строка s, возвращаем true если это палиндром, или false если иначе.
//

func main() {

	fmt.Println(isPalindrom("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrom("race a car"))                     // false
	fmt.Println(isPalindrom(" "))                              // true

}
