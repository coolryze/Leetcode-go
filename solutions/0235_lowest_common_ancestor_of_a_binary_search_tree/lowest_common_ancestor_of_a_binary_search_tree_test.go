package lowestcommonancestorofabinarysearchtree

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	root *TreeNode
	p    *TreeNode
	q    *TreeNode
}

type output struct {
	expected *TreeNode
}

func TestLowestCommonAncestor(t *testing.T) {
	root := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	tcs := []testcase{
		testcase{
			input{root, root.Left, root.Right},
			output{root},
		},
		testcase{
			input{root, root.Left, root.Left.Right},
			output{root.Left},
		},
	}

	for _, tc := range tcs {
		got := lowestCommonAncestor(tc.input.root, tc.input.p, tc.input.q)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
