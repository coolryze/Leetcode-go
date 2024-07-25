package mergeintervals

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	intervals [][]int
}

type output struct {
	expected [][]int
}

func TestMerge(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			output{[][]int{{1, 6}, {8, 10}, {15, 18}}},
		},
		testcase{
			input{[][]int{{1, 4}, {4, 5}}},
			output{[][]int{{1, 5}}},
		},
	}

	for _, tc := range tcs {
		got := merge(tc.input.intervals)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
