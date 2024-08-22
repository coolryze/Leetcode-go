package issubsequence

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
	expected bool
}

func TestIsSubsequence(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"abc", "ahbgdc"},
			output{true},
		},
		testcase{
			input{"axc", "ahbgdc"},
			output{false},
		},
	}

	for _, tc := range tcs {
		got := isSubsequence(tc.input.s, tc.input.t)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
