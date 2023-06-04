package main

import (
	"testing"
)

func TestLinerArraySearch(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{5, 7, 8}, 7, 1},
		{[]int{5, 7, 8}, 3, -1},
		{[]int{6, 6, 7}, 1, -1},
		{[]int{5, 7, 8, 12, 43, 54, 325, 765, 2, 6546, 3}, 3, 10},
		{[]int{6, 6, 7, 9, 0, 2, 3, 46, 77, 88, 3}, 1, -1},
	}
	for _, tc := range testCases {
		got := linearSearch(tc.nums, tc.target)
		if got != tc.want {
			t.Errorf("LinerArraySearch(%d, %d) = %d; want %d", tc.nums, tc.target, got, tc.want)
		}
	}
}
