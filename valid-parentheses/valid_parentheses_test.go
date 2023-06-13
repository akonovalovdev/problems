package main

import (
	"testing"
)

func TestValidParentheses(t *testing.T) {

	testCases := []struct {
		s    string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"())", false},
		{"((", false},
		{"(())))", false},
	}

	for _, tc := range testCases {
		got := isValid(tc.s)

		if got != tc.want {
			t.Errorf("Valid parentheses (%s) = %t, want: %t", tc.s, got, tc.want)
		}
	}
}
