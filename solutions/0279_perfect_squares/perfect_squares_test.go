package perfectsquares

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

func TestNumSquares(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{12},
			output{3},
		},
		testcase{
			input{13},
			output{2},
		},
	}

	for _, tc := range tcs {
		got := numSquares(tc.input.n)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
