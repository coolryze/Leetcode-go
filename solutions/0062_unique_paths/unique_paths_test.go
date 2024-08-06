package uniquepaths

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	m int
	n int
}

type output struct {
	expected int
}

func TestUniquePaths(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{3, 7},
			output{28},
		},
		testcase{
			input{3, 2},
			output{3},
		},
	}

	for _, tc := range tcs {
		got := uniquePaths(tc.input.m, tc.input.n)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
