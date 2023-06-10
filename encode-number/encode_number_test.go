package main

import "testing"

func TestEncodeNumber(t *testing.T) {
	testCases := []struct {
		input int
		want  string
	}{
		{123, "1111011"},
		{100, "1100100"},
		{0, ""},
		{100, "1100100"},
	}
	for _, tc := range testCases {
		got := encodeNumber(tc.input)

		if got != tc.want {
			t.Errorf("reverseString(%d) = %s; want %s", tc.input, got, tc.want)
		}
	}

}
