package largestrectangleinhistogram

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	height []int
}

type output struct {
	expected int
}

func TestLargestRectangleArea(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{2, 1, 5, 6, 2, 3}},
			output{10},
		},
		testcase{
			input{[]int{2, 4}},
			output{4},
		},
	}

	for _, tc := range tcs {
		got := largestRectangleArea(tc.input.height)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
