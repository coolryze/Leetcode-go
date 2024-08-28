package nextgreaterelementii

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
	expected []int
}

func TestNextGreaterElements(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 2, 1}},
			output{[]int{2, -1, 2}},
		},
		testcase{
			input{[]int{1, 2, 3, 4, 3}},
			output{[]int{2, 3, 4, -1, 4}},
		},
	}

	for _, tc := range tcs {
		got := nextGreaterElements(tc.input.nums)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
