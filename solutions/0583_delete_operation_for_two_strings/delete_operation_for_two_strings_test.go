package deleteoperationfortwostrings

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	word1 string
	word2 string
}

type output struct {
	expected int
}

func TestMinDistance(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"sea", "eat"},
			output{2},
		},
		testcase{
			input{"leetcode", "etco"},
			output{4},
		},
	}

	for _, tc := range tcs {
		got := minDistance(tc.input.word1, tc.input.word2)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
