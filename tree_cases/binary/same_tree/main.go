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
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Printf("%d, %d", p, q)
	//fmt.Println(isSameTree(p, j))
}

//
//func (n *TreeNode) Insert(value int) error {
//
//	if n == nil {
//		return errors.New("Cannot insert a value into a nil tree")
//	}
//
//	switch {
//	case value == n.Val:
//		return nil
//	case value < n.Val:
//		if n.Left == nil {
//			n.Left = &TreeNode{Val: value}
//			return nil
//		}
//		return n.Left.Insert(value)
//	case value > n.Val:
//		if n.Right == nil {
//			n.Right = &TreeNode{Val: value}
//			return nil
//		}
//		return n.Right.Insert(value)
//	}
//	return nil
//}
