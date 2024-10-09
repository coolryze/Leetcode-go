package repeatedsubstringpattern

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
}

type output struct {
	expected bool
}

func TestRepeatedSubstringPattern(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"abab"},
			output{true},
		},
		testcase{
			input{"aba"},
			output{false},
		},
		testcase{
			input{"abcabcabcabc"},
			output{true},
		},
	}

	for _, tc := range tcs {
		got := repeatedSubstringPattern(tc.input.s)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
