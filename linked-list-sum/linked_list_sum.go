package main

import "fmt"

type Slice interface {
	ToSlice() (digits []int)
}

type ListNode struct {
	Value int       // пустое значение - 0
	Next  *ListNode // пустое значение - nil
}

// шаг1. переводим срез в связный список
func NewListNode(digits []int) *ListNode {
	// это нужно потому(давайте предствим что этого нет) если мы вернём node,
	// и мы не сможем обратиться к другим элементам этого связного списка, так как это конец
	// то мы вернём последний элемент связного списка, а нам нужно вернуть
	// голову - указатель на начало связного списка
	head := &ListNode{}
	node := head
	for i := 0; i < len(digits); i++ {
		node.Next = &ListNode{
			Value: digits[i],
			Next:  nil,
		}
		fmt.Printf("%+v\n", node)
		node = node.Next
	}
	return head.Next
}

// шаг2. суммируем значения двух связных списков в 1 новый
func sumLinkedLists(a *ListNode, b *ListNode) *ListNode { //функция возвращает указательна на созданную структуру
	carry, dummy := 0, new(ListNode) //&ListNode{} мы создаём указатель(поинтер) на пустую структуру
	node := dummy
	for a != nil || b != nil || carry > 0 {
		if a != nil {
			carry += a.Value
			a = a.Next
		}
		if b != nil {
			carry += b.Value
			b = b.Next
		}
		node.Next = &ListNode{
			Value: carry % 10,
			Next:  nil,
		}
		node = node.Next
		carry /= 10
	}
	return dummy.Next
}

// шаг3. Переносим значение из  связного списка в новый слайс
func (node *ListNode) ToSlice() (digits []int) {
	for current := node; current != nil; current = current.Next {
		digits = append(digits, current.Value)
	}
	return digits
}

// (ДОПОЛНИТЕЛЬНЫЙ) шаг4. оптимизация вывода строки
func PrintSlice(slices ...Slice) {
	fmt.Println("Ниже работает интерфейс:")
	for _, oneSlice := range slices {
		nums := oneSlice.ToSlice()
		fmt.Println(nums)
	}
}
