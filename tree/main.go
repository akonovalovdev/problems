package main

// Есть сущность категория товаров, у нее есть айдишник и список дочерних (непосредственных прямых дочек):

type Category struct {
	ID       int
	Children []*Category
}

type CategoryTree struct {
	tree  []*Category
	cache map[int]struct{}
}

// уникальные категории помещать в мапу в качестве ключей, а значения пустые структуры
func (c *CategoryTree) Unique() []int {
	return uniqueIDs(c.tree, c.cache)
}

func uniqueIDs(categories []*Category, cache map[int]struct{}) []int {
	if len(categories) == 0 {
		return nil
	}
	var ids []int
	for _, c := range categories {
		_, ok := cache[c.ID]
		if !ok {
			ids = append(ids, c.ID)
			cache[c.ID] = struct{}{}
		}
		ids = append(ids, uniqueIDs(c.Children, cache)...)
	}

	return ids
}

// Тебе на вход даются корневые категории (массив деревьев): []*Category
//
// Сделай новый тип CategoriesTree у которого будут методы для:
//
// 1) получение числа уникальных категорий в дереве
// (теоретически, они могут повторяться - один и тот же ID в разных категориях)
//
// 2) получение максимальной глубины вложенности дерева
// (например, если всего 1 категория без детей - вложенность 1,
// если есть дети - 2, если хотя бы у одной из них есть дети - 3 и тд)
//
// 3) получение плоского списка категорий - то есть
// преобразование дерева в плоский список (когда родители и дети
// - на одном уровне массива - было {родитель->дети}, стало [родитель, ребенок1, ребенок2…]
