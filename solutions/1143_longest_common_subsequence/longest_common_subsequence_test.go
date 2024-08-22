package longestcommonsubsequence

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	text1 string
	text2 string
}

type output struct {
	expected int
}

func TestLongestCommonSubsequence(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"abcde", "ace"},
			output{3},
		},
		testcase{
			input{"abc", "abc"},
			output{3},
		},
		testcase{
			input{"abc", "def"},
			output{0},
		},
	}

	for _, tc := range tcs {
		got := longestCommonSubsequence(tc.input.text1, tc.input.text2)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
