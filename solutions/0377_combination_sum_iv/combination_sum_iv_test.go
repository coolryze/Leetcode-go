package combinationsumiv

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

func TestCombinationSum4(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 2, 3}, 4},
			output{7},
		},
		testcase{
			input{[]int{9}, 3},
			output{0},
		},
	}

	for _, tc := range tcs {
		got := combinationSum4(tc.input.nums, tc.input.target)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
