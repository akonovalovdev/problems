package main

import "testing"

func TestSameTree(t *testing.T) {
	testCases := []struct {
		p    *TreeNode
		q    *TreeNode
		want bool
	}{
		//	Input: p = [1,2,3], q = [1,2,3]
		//	Output: true
		{&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			true,
		},
		//	Input: p = [1,2], q = [1,null,2]
		// 	Output: false
		{&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
		},
			&TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			false,
		},
	}

	for _, tc := range testCases {
		got := isSameTree(tc.p, tc.q)

		if tc.want != got {
			t.Errorf("SameTree(%v, %v) = %t, want = %t", tc.p, tc.q, got, tc.want)
		}

	}

}
