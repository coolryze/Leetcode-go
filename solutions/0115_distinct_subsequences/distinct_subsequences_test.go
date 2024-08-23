package distinctsubsequences

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	s string
	t string
}

type output struct {
	expected int
}

func TestNumDistinct(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"rabbbit", "rabbit"},
			output{3},
		},
		testcase{
			input{"babgbag", "bag"},
			output{5},
		},
	}

	for _, tc := range tcs {
		got := numDistinct(tc.input.s, tc.input.t)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
