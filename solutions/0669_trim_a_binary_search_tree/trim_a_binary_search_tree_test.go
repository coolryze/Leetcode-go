package trimabinarysearchtree

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	root  []interface{}
	low   int
	hight int
}

type output struct {
	expected []interface{}
}

func arrayToTree(nums []interface{}) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0].(int)}
	queue := []*TreeNode{root}
	index := 1

	for index < n {
		node := queue[0]
		queue = queue[1:]

		if index < n && nums[index] != nil {
			left := &TreeNode{Val: nums[index].(int)}
			node.Left = left
			queue = append(queue, left)
		}
		index++

		if index < n && nums[index] != nil {
			right := &TreeNode{Val: nums[index].(int)}
			node.Right = right
			queue = append(queue, right)
		}
		index++
	}

	return root
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
			input{[]interface{}{1, 0, 2}, 1, 2},
			output{[]interface{}{1, nil, 2}},
		},
		testcase{
			input{[]interface{}{1, 0, 2}, 1, 2},
			output{[]interface{}{1, nil, 2}},
		},
		testcase{
			input{[]interface{}{3, 0, 4, nil, 2, nil, nil, 1}, 1, 3},
			output{[]interface{}{3, 2, nil, 1}},
		},
	}

	for _, tc := range tcs {
		got := treeToLevelOrderArray(trimBST(arrayToTree(tc.input.root), tc.input.low, tc.input.hight))
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
