package palindromicsubstrings

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

func TestCountSubstrings(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"abc"},
			output{3},
		},
		testcase{
			input{"aaa"},
			output{6},
		}, testcase{
			input{"cbabc"},
			output{7},
		},
	}

	for _, tc := range tcs {
		got := countSubstrings(tc.input.s)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
