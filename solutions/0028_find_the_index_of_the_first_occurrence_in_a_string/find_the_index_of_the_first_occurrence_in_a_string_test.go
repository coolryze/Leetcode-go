package findtheindexofthefirstoccurrenceinastring

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	haystack string
	needle   string
}

type output struct {
	expected int
}

func TestStrStr(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"sadbutsad", "sad"},
			output{0},
		},
		testcase{
			input{"leetcode", "leeto"},
			output{-1},
		},
		testcase{
			input{"mississippi", "issip"},
			output{4},
		},
	}

	for _, tc := range tcs {
		got := strStr(tc.input.haystack, tc.input.needle)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
