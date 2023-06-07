package main

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{"Привет я Петя!", "!ятеП я тевирП"},
		{"Мой номер телефона: 89167854141", "14145876198 :анофелет ремон йоМ"},
		{"Hello, 世界", "界世 ,olleH"},
	}
	for _, tc := range testCases {
		got := reverseString(tc.input)

		if got != tc.want {
			t.Errorf("reverseString(%s) = %s; want %s", tc.input, got, tc.want)
		}
	}

}
