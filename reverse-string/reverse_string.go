package main

func reverseString(input string) string {
	// принимаемую строку переделываем в слайс рун для внесения корректировок
	arr := []rune(input)
	for i := 0; i < len(arr)/2; i++ {
		end := len(arr) - i - 1
		// без помощи "третьей полки" переставляем первый с последним символом местами, сдвигаясь к центру на 1
		arr[i] = arr[i] + arr[end]
		arr[end] = arr[i] - arr[end]
		arr[i] = arr[i] - arr[end]
		// либо можно arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]]
	}
	//возвращаем переставленный слайс рун, преобразовывая его в строку
	return string(arr)
}
