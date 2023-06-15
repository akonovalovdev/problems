package main

import "fmt"

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
