package minimumnumberofarrowstoburstballoons

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	points [][]int
}

type output struct {
	expected int
}

func TestFindMinArrowShots(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}},
			output{2},
		},
		testcase{
			input{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}},
			output{4},
		},
		testcase{
			input{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}},
			output{2},
		},
	}

	for _, tc := range tcs {
		got := findMinArrowShots(tc.input.points)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
