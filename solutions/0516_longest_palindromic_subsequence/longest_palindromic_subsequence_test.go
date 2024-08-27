package longestpalindromicsubsequence

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
	expected int
}

func TestLongestPalindromeSubseq(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"bbbab"},
			output{4},
		},
		testcase{
			input{"cbbd"},
			output{2},
		},
	}

	for _, tc := range tcs {
		got := longestPalindromeSubseq(tc.input.s)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
