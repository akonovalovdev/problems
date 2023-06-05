package main

// Дано цело число `day`. Нужно напечатать в консоль день недели,
//соответствующий данному целому числу: 1 → Понедельник, 2 → Вторник, …, 7 → Воскресенье.
//
// Подсказки (даю подсказки, потому что только начало, потом подсказок не будет):
//
// - Целое число может быть как 0, так и отрицательным, так и больше 7
// - В случае “пограничных случаев” (edge cases) - выводить в консоль “Некорректный день недели”

import "fmt"

func main() {
	fmt.Println(weekday(5))
	fmt.Println(weekday(1))
	fmt.Println(weekday(10))
	fmt.Println(weekday(-5))
}