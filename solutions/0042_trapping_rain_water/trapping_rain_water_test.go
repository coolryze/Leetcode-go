package trappingrainwater

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

func TestTrap(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}},
			output{6},
		},
		testcase{
			input{[]int{4, 2, 0, 3, 2, 5}},
			output{9},
		},
	}

	for _, tc := range tcs {
		got := trap(tc.input.height)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
