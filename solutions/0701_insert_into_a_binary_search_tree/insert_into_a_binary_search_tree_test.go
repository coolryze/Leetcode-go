package insertintoabinarysearchtree

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	nums   []interface{}
	target int
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

func TestInsertIntoBST(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]interface{}{4, 2, 7, 1, 3}, 5},
			output{[]interface{}{4, 2, 7, 1, 3, 5}},
		},
		testcase{
			input{[]interface{}{40, 20, 60, 10, 30, 50, 70}, 25},
			output{[]interface{}{40, 20, 60, 10, 30, 50, 70, nil, nil, 25}},
		},
		testcase{
			input{[]interface{}{4, 2, 7, 1, 3, nil, nil, nil, nil, nil, nil}, 5},
			output{[]interface{}{4, 2, 7, 1, 3, 5}},
		},
	}

	for _, tc := range tcs {
		got := treeToLevelOrderArray(insertIntoBST(arrayToTree(tc.input.nums), tc.input.target))
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
