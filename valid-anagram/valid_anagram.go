package main

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	anagr := make(map[rune]int, 33)
	for _, v := range s {
		anagr[v] = anagr[v] + 1
	}
	for _, v := range t {
		ostatok, ok := anagr[v]
		if !ok {
			return false
		}
		anagr[v] = ostatok - 1
		if ostatok == 1 { // вот так корректно считает
			delete(anagr, v)
		}
	}
	return true
}
