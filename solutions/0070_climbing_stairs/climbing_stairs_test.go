package climbingstairs

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	n int
}

type output struct {
	expected int
}

func TestClimbStairs(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{2},
			output{2},
		},
		testcase{
			input{3},
			output{3},
		},
	}

	for _, tc := range tcs {
		got := climbStairs(tc.input.n)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
