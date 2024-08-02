package mincostclimbingstairs

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	cost []int
}

type output struct {
	expected int
}

func TestMinCostClimbingStairs(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{10, 15, 20}},
			output{15},
		},
		testcase{
			input{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}},
			output{6},
		},
	}

	for _, tc := range tcs {
		got := minCostClimbingStairs(tc.input.cost)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
