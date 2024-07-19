package gasstation

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	gas  []int
	cost []int
}

type output struct {
	expected int
}

func TestCanCompleteCircuit(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}},
			output{3},
		},
		testcase{
			input{[]int{2, 3, 4}, []int{3, 4, 3}},
			output{-1},
		},
	}

	for _, tc := range tcs {
		got := canCompleteCircuit(tc.input.gas, tc.input.cost)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
