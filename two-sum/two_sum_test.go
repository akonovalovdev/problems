package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{1, 2, 3}, 5, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{9, 5, 4, 3, 8, 7}, 9, []int{1, 2}},
		{[]int{9, 5, 1, 4, 3, 8, 7, 3, 0, 5, 1}, 2, []int{2, 10}},
		{[]int{91, 53, 42, 32, 82, 7, 8, 45, 95}, 99, []int{0, 6}},
	}
	for _, tc := range testCases {
		got := twoSum(tc.nums, tc.target)

		if !equalSlices(got, tc.want) {
			t.Errorf("twoSum(%d,%d) = %d; want %d", tc.nums, tc.target, got, tc.want)
		}
	}

}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
