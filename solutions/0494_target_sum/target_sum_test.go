package targetsum

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	nums   []int
	target int
}

type output struct {
	expected int
}

func TestFindTargetSumWays(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 1, 1, 1, 1}, 3},
			output{5},
		},
		testcase{
			input{[]int{1}, 1},
			output{1},
		},
	}

	for _, tc := range tcs {
		got := findTargetSumWays(tc.input.nums, tc.input.target)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
