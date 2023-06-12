package main

import "fmt"

// 1-убрать пробелы
// сравнить первый элемент с последним, если они не равны, то вернуть false, если цикл завершен то вернуть true
func isPalindrom(input string) bool {
	sli := []rune(input)
	var nSli []rune
	for i := 0; i < len(input); i++ {
		if sli[i] < 'a' {
			sli[i] = sli[i] + 'a' - 'A'
		}
		if sli[i] < 'a' || sli[i] > 'z' {
			continue
		}
		nSli = append(nSli, sli[i])
	}

	for i := 0; i < len(nSli)/2; i++ {
		if nSli[i] != nSli[len(nSli)-1-i] {
			fmt.Printf("Объяснение: %q - не палиндром.\n", string(nSli))
			return false
		}
	}
	fmt.Printf("Объяснение: %q - палиндром.\n", string(nSli))
	return true
}
