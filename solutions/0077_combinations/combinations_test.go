package combinations

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	n int
	k int
}

type output struct {
	expected [][]int
}

func TestCombine(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{4, 2},
			output{[][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		},
		testcase{
			input{1, 1},
			output{[][]int{{1}}},
		},
	}

	for _, tc := range tcs {
		got := combine(tc.input.n, tc.input.k)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
