package deletenodeinabst

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

func TestDeleteNode(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]interface{}{5, 3, 6, 2, 4, nil, 7}, 3},
			output{[]interface{}{5, 4, 6, 2, nil, nil, 7}},
		},
		testcase{
			input{[]interface{}{5, 3, 6, 2, 4, nil, 7}, 0},
			output{[]interface{}{5, 3, 6, 2, 4, nil, 7}},
		},
		testcase{
			input{[]interface{}{}, 0},
			output{[]interface{}{}},
		},
	}

	for _, tc := range tcs {
		got := treeToLevelOrderArray(deleteNode(arrayToTree(tc.input.nums), tc.input.target))
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
