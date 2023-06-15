package main

import "fmt"

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil { // исключаем панику в случаях, когда одно или несколько деревьев на входе пришли со значением nil
		return q == nil // если оба nil, то вернём true
	} else if q == nil {
		return false
	}
	if p.Val != q.Val { // сравниваем значения одинаково расположенных ветвей
		fmt.Printf("значения p = %d, q = %d - не совпали\n", p.Val, q.Val)
		return false
	}
	fmt.Printf("значения p = %d, q = %d - совпали\n", p.Val, q.Val)
	// проверяем симметрию наличия либо отсутствия ветвей-детей слева
	if p.Left != nil && q.Left != nil || p.Left == nil && q.Left == nil {
		fmt.Println("Наличие левых детей, одинаковое")
	} else {
		fmt.Println("Наличие левых детей, не одинаковое")
		return false
	}
	// проверяем симметрию наличия либо отсутствия ветвей-детей слева
	if p.Right != nil && q.Right != nil || p.Right == nil && q.Right == nil {
		fmt.Println("Наличие правых детей, одинаковое")
	} else {
		fmt.Println("Наличие правых детей, не одинаковое")
		return false
	}
	//проверяем значения детей-ветвей слева, если они существуют
	if p.Left != nil {
		fmt.Println("Проверяем левых детей, так как они существуют")
		otvet := isSameTree(p.Left, q.Left)
		if otvet != true {
			return false
		}
	}

	//проверяем значения детей-ветвей справа, если они существуют
	if p.Right != nil {
		fmt.Println("Проверяем правых детей, так как они существуют")
		otvet := isSameTree(p.Right, q.Right)
		if otvet != true {
			return false
		}
	}
	//все проверки пройдены, деревья одинаковые
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
