package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	testCases := []struct {
		input string
		want  bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}

	for _, tc := range testCases {
		got := isPalindrome(tc.input)

		if got != tc.want {
			t.Errorf("Is palindrome(%s) = %t, want %t", tc.input, got, tc.want)
		}
	}
}
