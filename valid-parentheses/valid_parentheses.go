package main

import "fmt"

// 1. создаём второй массив(стэк), куда складываем все открывающие скобки
// 2. пробегаемся по исходному массиву циклом, смотрим тип скобок, открытые или закрытые
// 3. соответствующие закрывающие удаляем, при встрече не соответствующих возвращаем false
func isValid(s string) bool {
	sli := []rune(s)
	nSli := make([]rune, 0, len(sli)/2)
	if len(sli)%2 != 0 {
		fmt.Printf("Один лишний элемент в строке: ")
		return false
	}
	for i := 0; i < len(sli); i++ {
		if sli[i] == '[' || sli[i] == '{' || sli[i] == '(' {
			nSli = append(nSli, sli[i])
			continue
		}

		switch {
		case len(nSli) == 0:
			fmt.Printf("В стэке нет скобок: ")
			return false
		case sli[i] == ')' && ')'-1 == nSli[len(nSli)-1]:
			nSli = nSli[:len(nSli)-1]
		case sli[i]-2 == nSli[len(nSli)-1]:
			nSli = nSli[:len(nSli)-1]
		default:
			fmt.Printf("Скобки не совпали: ")
			return false
		}
	}
	if len(nSli) != 0 {
		fmt.Printf("Не все скобки были закрыты: ")
		return false
	}
	return true
}
