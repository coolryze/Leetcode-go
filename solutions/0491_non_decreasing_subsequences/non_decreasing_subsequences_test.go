package nondecreasingsubsequences

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
	expected [][]int
}

func TestFindSubsequences(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{4, 6, 7, 7}},
			output{[][]int{{4, 6}, {4, 6, 7}, {4, 6, 7, 7}, {4, 7}, {4, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}}},
		},
		testcase{
			input{[]int{4, 4, 3, 2, 1}},
			output{[][]int{{4, 4}}},
		},
	}

	for _, tc := range tcs {
		got := findSubsequences(tc.input.nums)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
