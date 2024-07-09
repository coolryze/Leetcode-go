package convertsortedarraytobinarysearchtree

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	nums []int
}

type output struct {
	expected []interface{}
}

func treeToLevelOrderArray(root *TreeNode) []interface{} {
	res := []interface{}{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			res = append(res, node.Val)
			queue = append(queue, node.Left, node.Right)
		} else {
			res = append(res, nil)
		}
	}

	for len(res) > 0 && res[len(res)-1] == nil {
		res = res[:len(res)-1]
	}

	return res
}

func TestTrimBST(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{-10, -3, 0, 5, 9}},
			output{[]interface{}{0, -3, 9, -10, nil, 5}},
		},
		testcase{
			input{[]int{1, 3}},
			output{[]interface{}{3, 1}},
		},
	}

	for _, tc := range tcs {
		got := treeToLevelOrderArray(sortedArrayToBST(tc.input.nums))
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
