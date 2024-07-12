package palindromepartitioning

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
	expected [][]string
}

func TestPartition(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"aab"},
			output{[][]string{{"a", "a", "b"}, {"aa", "b"}}},
		},
		testcase{
			input{"a"},
			output{[][]string{{"a"}}},
		},
	}

	for _, tc := range tcs {
		got := partition(tc.input.s)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
