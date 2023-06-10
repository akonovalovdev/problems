package main

import "fmt"

func encodeNumber(input int) (binary string) {
	for input >= 1 {
		ostatok := input % 2
		binary = fmt.Sprintf("%d%s", ostatok, binary)
		input /= 2
	}
	return binary
}

// 123/2=61 % 1
// 61/2=30 % 1
// 30/2=15 % 0
// 15/2=7 % 1
// 7/2=3 % 1
// 3/2=1 % 1
// 1/2 % 1
