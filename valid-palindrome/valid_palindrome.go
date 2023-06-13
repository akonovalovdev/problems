package main

import (
	"fmt"
)

// 1-убрать пробелы
// сравнить первый элемент с последним, если они не равны, то вернуть false, если цикл завершен то вернуть true
func isPalindrome(input string) bool {
	// для изменения текста, переводим строку в слайс рун
	sli := []rune(input)
	// создаём дополнительный слайс, куда будем записывать только нужные символы
	var nSli []rune
	for i := 0; i < len(input); i++ {
		// переводим все английские буквы из верхнего регистра в нижний
		if sli[i] >= 'A' && sli[i] <= 'Z' {
			sli[i] = sli[i] + 'a' - 'A'
		}
		// исключаем все символы кроме английских строчных букв
		if sli[i] < '0' || sli[i] > '9' && sli[i] < 'a' || sli[i] > 'z' {
			continue
		}
		// добавляем соответсвующий требованиям символ в новый слайс
		nSli = append(nSli, sli[i])
	}

	for i := 0; i < len(nSli)/2; i++ {
		// сверяем зеркальное соответствие начала с концом получившегося слайса
		if nSli[i] != nSli[len(nSli)-1-i] {
			fmt.Printf("Объяснение: %q - не палиндром.\n", string(nSli))
			return false
		}
	}
	fmt.Printf("Объяснение: %q - палиндром.\n", string(nSli))
	return true
}

// альтернативное решение с методом пакета unicode .ToLover
//func isPalindrome(s string) bool {
//	i := 0
//	j := len(s) - 1
//	arr := []rune(s)
//
//	for i < j {
//		left := unicode.ToLower(arr[i])
//		right := unicode.ToLower(arr[j])
//
//		if !isLetterOrDigit(left) {
//			i++
//			continue
//		}
//
//		if !isLetterOrDigit(right) {
//			j--
//			continue
//		}
//
//		if left != right {
//			return false
//		}
//
//		i++
//		j--
//	}
//
//	return true
//}
//
//
//func isLetterOrDigit(s rune) bool {
//	return unicode.IsLetter(s) || unicode.IsDigit(s)
//}
