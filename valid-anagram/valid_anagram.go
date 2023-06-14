package main

func isAnagram(s, t string) bool {
	// исключаем кейсы, когда длины обеих строк не равны
	if len(s) != len(t) {
		return false
	}
	// создаём мапу в ключах которой будем хранить все имеющиеся буквы первого слова, а значении их повторы
	anagr := make(map[rune]int)
	for _, v := range s { // добавляем ключ и значение по ключу
		anagr[v] = anagr[v] + 1 // прибавляем счётчик повторений буквы
	}
	// пробегаемся циклом по второй сроке, находя и удалял ключи из мапы
	for _, v := range t {
		ostatok, ok := anagr[v]
		if !ok { // если не находим искомы ключ, возвращаем ложь
			return false
		}
		anagr[v] = ostatok - 1 // уменьшаем счётчик повторений
		ostatok = anagr[v]
		// удаляем ключ, если остаток стал нулевым, если оставить тогда может начать уходить в минус
		if ostatok == 0 {
			delete(anagr, v)
		}
	}
	return true
}

// альтернативный вариант решения
//func isAnagram(s string, t string) bool {
//	if len(s) != len(t) {
//		return false
//	}
//
//	var freq [26]int // создаём на 26 мест
//
//	for idx := 0; idx < len(s); idx++ { // отнимаем и добавляем разницу символов в байтах
//		freq[s[idx] - 'a']++
//		freq[t[idx] - 'a']--
//	}
//
//	for idx := 0; idx < len(freq); idx++ { // проверяем что все элементы равны нулю
//		if freq[idx] != 0 {
//			return false
//		}
//	}
