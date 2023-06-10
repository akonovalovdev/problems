package main

import "fmt"

func encodeNumber(input int) (binary string) {
	// в цикле переводим заданное число в бинарные значение
	for input >= 1 {
		ostatok := input % 2                          // каждую итерацию берём остаток 2 от числа
		binary = fmt.Sprintf("%d%s", ostatok, binary) // подставляем остаток в строку слева от записанных ранее
		input /= 2                                    // уменьшаем число на 2
	}
	return binary // возвращаем строку как ответ
}

// 123/2=61 % 1
// 61/2=30 % 1
// 30/2=15 % 0
// 15/2=7 % 1
// 7/2=3 % 1
// 3/2=1 % 1
// 1/2 % 1
