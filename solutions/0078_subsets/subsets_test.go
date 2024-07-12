package subsets

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

func TestSubsets(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 2, 3}},
			output{[][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}},
		},
		testcase{
			input{[]int{0}},
			output{[][]int{{}, {0}}},
		},
	}

	for _, tc := range tcs {
		got := subsets(tc.input.nums)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
