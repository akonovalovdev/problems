package main

// Вам даны два непустых **односвязных списка**, представляющих два неотрицательных целых числа.
// Цифры хранятся в обратном порядке, и каждый из их узлов содержит одну цифру.
// Просуммируйте два числа и верните сумму в виде связанного списка.
//
// Вы можете считать, что эти два числа не содержат начальных нулей, если само по себе число не равно 0.
//
// # Ограничения
// - Количество узлов в каждом связанном списке находится в диапазоне [1, 100]
// - Значение в каждой ноде - число от 0 до 9: 0 <= Node.Value <= 9
// - Гарантируется, что список представляет собой число, не имеющее лидирующих нулей:
// не будет чисел, представленных в виде 000123 — будет 123.
// При этом число 123 будет записано в форме цепочки: 3→2→1.

import "fmt"

func main() {

	a := NewListNode([]int{2, 4, 3})
	b := NewListNode([]int{5, 6, 4}) //342+465
	sum := sumLinkedLists(a, b)
	fmt.Println(sum)
	fmt.Println(sum.ToSlice())
	PrintSlice(a, b, sum) // дополнительный шаг
}
