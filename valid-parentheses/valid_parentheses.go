package main

import "fmt"

// 1. создаём второй массив(стэк), куда складываем все открывающие скобки
// 2. пробегаемся по исходному массиву циклом, смотрим тип скобок, открытые или закрытые
// 3. соответствующие закрывающие удаляем, при встрече не соответствующих возвращаем false
func isValid(s string) bool {
	input := []rune(s)                     // переводим строку в срез рун
	stack := make([]rune, 0, len(input)/2) // создаём срез-стэк
	// исключаем затраты ресурсов на работу с кейсами с нечетным количеством элементов
	if len(input)%2 != 0 {
		fmt.Printf("Один лишний элемент в строке: ")
		return false
	}
	// пробегаемся по входному срезу рун
	for i := 0; i < len(input); i++ {
		// все открывающие скобки записываем в срез-стэк
		if input[i] == '[' || input[i] == '{' || input[i] == '(' {
			stack = append(stack, input[i])
			continue // завершаем итерацию цикла
		}
		// пробегаемся по всем остальным видам скобок и находим нужную
		if len(stack) == 0 { // исключающий панику кес с пустым стеком(либо не добавили ещё, либо уже удалили)
			fmt.Printf("В стэке нет скобок: ")
			return false
		}
		if input[i] == ')' && '(' == stack[len(stack)-1] {
			stack = stack[:len(stack)-1] // удаляем совпавшую скобку
		}
		if input[i] == '}' && '{' == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if input[i] == ']' && '[' == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
	}
	// исключаем кейс когда в срезе-стэке всё-ещё остались незакрытые скобки, а в основном все скобки закончились
	// победа, цикл завершен

	if len(stack) != 0 {
		fmt.Printf("Скобки не совпали: ")
		return false
	}

	return true
}

// альтернативное решение

//func isValid(s string) bool {
//	pairs := map[byte]byte{ // создаётся мапа, где ключ закрывающая скобка, а значение открывающая
//		'}': '{',
//		']': '[',
//		')': '(',
//	}
//
//	stack := make([]byte, 0)
//
//	for _, char := range []byte(s) {
//		pair, ok := pairs[char]
//		if !ok {
//			stack = append(stack, char)
//			continue
//		}
//
//		if len(stack) == 0 {
//			return false
//		}
//
//		if stack[len(stack) - 1] != pair {
//			return false
//		}
//
//		stack = stack[:len(stack) - 1]
//	}
//
//	return len(stack) == 0
//}
