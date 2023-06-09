package main

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {

	testCases := []struct {
		s    string
		t    string
		want bool
	}{
		{"анаграмма", "нагарамма", true},
		{"крыса", "автомобиль", false},
		{"истребитель", "бистретеиль", true},
	}
	for _, tc := range testCases {
		got := isAnagram(tc.s, tc.t)

		if got != tc.want {
			t.Errorf("Is anagram(%s, %s) = %t, want %t", tc.s, tc.t, got, tc.want)
		}
	}
}
