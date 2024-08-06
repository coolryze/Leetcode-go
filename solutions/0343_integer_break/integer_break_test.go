package integerbreak

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

func TestIntegerBreak(t *testing.T) {
	tcs := []testcase{

		testcase{
			input{2},
			output{1},
		},
		testcase{
			input{10},
			output{36},
		},
	}

	for _, tc := range tcs {
		got := integerBreak(tc.input.n)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
