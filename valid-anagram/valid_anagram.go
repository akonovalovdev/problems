package main

func isAnagram(s, t string) bool {
	anagr := make(map[rune]int)
	for _, v := range s {
		anagr[v] = anagr[v] + 1
	}
	for _, v := range t {
		ostatok, ok := anagr[v]
		if !ok {
			return false
		}
		anagr[v] = ostatok - 1
		ostatok = anagr[v]
		if ostatok == 0 {
			delete(anagr, v)
		}
	}
	return true
}
