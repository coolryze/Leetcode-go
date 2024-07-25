package binarytreecameras

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	root []interface{}
}

type output struct {
	expected int
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

func TestMinCameraCover(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]interface{}{0, 0, nil, 0, 0}},
			output{1},
		},
		testcase{
			input{[]interface{}{0, 0, nil, 0, nil, 0, nil, nil, 0}},
			output{2},
		},
	}

	for _, tc := range tcs {
		got := minCameraCover(arrayToTree(tc.input.root))
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
