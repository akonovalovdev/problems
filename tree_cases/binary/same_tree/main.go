package main

import (
	"fmt"
)

// Имея корни двух бинарных деревьев p и q, напишите функцию, проверяющую, совпадают ли они или нет.
//
// Два бинарных дерева считаются одинаковыми, если они структурно идентичны, а узлы имеют одинаковое значение.

// Объявление узла двоичного дерева
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//Вход: p = [1,2,3], q = [1,2,3]
	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(isSameTree(p, q))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil {
		return q == nil
	} else if q == nil {
		return false
	}
	if p.Val != q.Val {
		fmt.Printf("значения p = %d, q = %d - не совпали\n", p.Val, q.Val)
		return false
	}
	fmt.Printf("значения p = %d, q = %d - совпали\n", p.Val, q.Val)
	if p.Left != nil && q.Left != nil || p.Left == nil && q.Left == nil {
		fmt.Println("Наличие левых детей, одинаковое")
	} else {
		fmt.Println("Наличие левых детей, не одинаковое")
		return false
	}
	if p.Right != nil && q.Right != nil || p.Right == nil && q.Right == nil {
		fmt.Println("Наличие правых детей, одинаковое")
	} else {
		fmt.Println("Наличие правых детей, не одинаковое")
		return false
	}

	if p.Left != nil {
		fmt.Println("Проверяем левых детей, так как они существуют")
		otvet := isSameTree(p.Left, q.Left)
		if otvet != true {
			return false
		}
	}
	if p.Right != nil {
		fmt.Println("Проверяем правых детей, так как они существуют")
		otvet := isSameTree(p.Right, q.Right)
		if otvet != true {
			return false
		}
	}

	return true
}
