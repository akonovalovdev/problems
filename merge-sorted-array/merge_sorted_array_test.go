package main

import "testing"

func TestMerge(t *testing.T) {
	testCases := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, 1, []int{}, 0, []int{1}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
		{[]int{2, 0}, 1, []int{1}, 1, []int{1, 2}},
		{[]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 4, 5, 6}},
		{[]int{-1, 0, 0, 3, 3, 3, 0, 0, 0}, 6, []int{1, 2, 2}, 3, []int{-1, 0, 0, 1, 2, 2, 3, 3, 3}},
		{[]int{-1, 0, 0, 3, 3, 3, 0, 0, 0}, 6, []int{-1, 1, 2}, 3, []int{-1, -1, 0, 0, 1, 2, 3, 3, 3}},
		{[]int{0, 0, 0, 0, 0}, 0, []int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{-1, 1, 0, 0, 0, 0, 0, 0}, 2, []int{-1, 0, 1, 1, 2, 3}, 6, []int{-1, -1, 0, 1, 1, 1, 2, 3}},
		{[]int{-1, 0, 0, 3, 3, 3, 0, 0, 0, 0}, 6, []int{-1, 0, 1, 4}, 4, []int{-1, -1, 0, 0, 0, 1, 3, 3, 3, 4}},
	}

	for _, tc := range testCases {
		merge(tc.nums1, tc.m, tc.nums2, tc.n)
		got := tc.nums1
		if !equalSlices(got, tc.want) {
			t.Errorf("MergeSortedArray(%d, %d, %d, %d) = %d, want %d", tc.nums1, tc.m, tc.nums2, tc.n, got, tc.want)
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
