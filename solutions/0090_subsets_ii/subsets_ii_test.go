package subsetsii

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

func TestSubsetsWithDup(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 2, 2}},
			output{[][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
		},
		testcase{
			input{[]int{0}},
			output{[][]int{{}, {0}}},
		},
	}

	for _, tc := range tcs {
		got := subsetsWithDup(tc.input.nums)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
