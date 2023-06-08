package main

import (
	"testing"
)

func TestConstainsDuplicate(t *testing.T) {
	testCases := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{[]int{1, 2, 3, 4}, false},
	}
	for _, tc := range testCases {
		got := constainsDuplicate(tc.nums)

		if got != tc.want {
			t.Errorf("constainsDuplicate(%d) = %t, want %t", tc.nums, got, tc.want)
		}
	}
}
