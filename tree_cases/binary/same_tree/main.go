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
