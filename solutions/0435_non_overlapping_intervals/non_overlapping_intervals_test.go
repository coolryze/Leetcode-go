package nonoverlappingintervals

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
	expected int
}

func TestEraseOverlapIntervals(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}},
			output{1},
		},
		testcase{
			input{[][]int{{1, 2}, {1, 2}, {1, 2}}},
			output{2},
		},
		testcase{
			input{[][]int{{1, 2}, {2, 3}}},
			output{0},
		},
		testcase{
			input{[][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}}},
			output{2},
		},
	}

	for _, tc := range tcs {
		got := eraseOverlapIntervals(tc.input.intervals)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
