package laststoneweightii

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	stones []int
}

type output struct {
	expected int
}

func TestLastStoneWeightII(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{2, 7, 4, 1, 8, 1}},
			output{1},
		},
		testcase{
			input{[]int{31, 26, 33, 21, 40}},
			output{5},
		},
	}

	for _, tc := range tcs {
		got := lastStoneWeightII(tc.input.stones)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
