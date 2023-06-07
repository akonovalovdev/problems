package main

func reverseString(input string) string {
	arr := []rune(input)
	for i := 0; i < len(arr)/2; i++ {
		end := len(arr) - i - 1 // либо можно arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]]
		arr[i] = arr[i] + arr[end]
		arr[end] = arr[i] - arr[end]
		arr[i] = arr[i] - arr[end]
	}
	return string(arr)
}
