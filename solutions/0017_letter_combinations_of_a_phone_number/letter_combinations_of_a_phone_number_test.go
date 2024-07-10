package lettercombinationsofaphonenumber

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	digits string
}

type output struct {
	expected []string
}

func TestLetterCombinations(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"23"},
			output{[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		},
		testcase{
			input{""},
			output{[]string{}},
		},
		testcase{
			input{"2"},
			output{[]string{"a", "b", "c"}},
		},
	}

	for _, tc := range tcs {
		got := letterCombinations(tc.input.digits)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
