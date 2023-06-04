package main

import (
	"testing"
)

func TestBinSearch(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5, 7, 9, 10, 11}, 10, 7},
		{[]int{1, 2, 3, 4, 5, 7, 9}, 5, 4},
		{[]int{3, 4, 5, 7, 9}, 9, 4},
		{[]int{5, 7, 9, 11}, 11, 3},
		{[]int{1, 2, 3, 4, 5, 7, 9, 10, 11}, 0, -1},
		{[]int{1, 2, 3, 4, 5, 7, 9}, -100, -1},
		{[]int{3, 4, 5, 7, 9}, 100, -1},
		{[]int{5, 7, 9, 11}, 8, -1},
	}
	for _, tc := range testCases {
		got := binSearch(tc.nums, tc.target)
		if got != tc.want {
			t.Errorf("BinSearch(%d, %d) = %d; want %d", tc.nums, tc.target, got, tc.want)
		}
	}
}
