package main

import (
	"testing"
)

func TestSumLinkedLists(t *testing.T) {
	// Test cases
	tests := []struct {
		a        []int
		b        []int
		expected []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{[]int{0}, []int{0}, []int{0}},
		{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}

	// Test each case
	for _, test := range tests {
		a := NewListNode(test.a)
		b := NewListNode(test.b)
		expected := NewListNode(test.expected)

		result := sumLinkedLists(a, b)

		// Compare each node of the linked lists
		for expected != nil && result != nil {
			if expected.Value != result.Value {
				t.Errorf("Sum of %v and %v was incorrect, got: %v, want: %v.", test.a, test.b, result.ToSlice(), test.expected)
			}

			expected = expected.Next
			result = result.Next
		}

		// Check if the lengths of the linked lists match
		if expected != result {
			t.Errorf("Sum of %v and %v was incorrect, got: %v, want: %v.", test.a, test.b, result.ToSlice(), test.expected)
		}
	}
}
