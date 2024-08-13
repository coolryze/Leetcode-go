package wordbreak

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	s        string
	wordDict []string
}

type output struct {
	expected bool
}

func TestWordBreak(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"leetcode", []string{"leet", "code"}},
			output{true},
		},
		testcase{
			input{"applepenapple", []string{"apple", "pen"}},
			output{true},
		},
		testcase{
			input{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}},
			output{false},
		},
	}

	for _, tc := range tcs {
		got := wordBreak(tc.input.s, tc.input.wordDict)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
