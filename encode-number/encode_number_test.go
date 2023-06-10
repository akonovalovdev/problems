package main

import "testing"

func TestFlifArraySearch(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{[]int{5, 7, 8}, []int{8, 7, 5}},
		{[]int{2, 5, 3, 8}, []int{8, 3, 5, 2}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, []int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{[]int{1}, []int{1}},
	}
	for _, tc := range testCases {
		got := flipArray(tc.nums)
		if !equalSlices(got, tc.want) {
			t.Errorf("flipArraySearch(%d) = %d, want %d", tc.nums, got, tc.want)
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
