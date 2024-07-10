package combinationsumiii

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

func TestCombinationSum3(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{3, 7},
			output{[][]int{{1, 2, 4}}},
		},
		testcase{
			input{3, 9},
			output{[][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		},
		testcase{
			input{4, 1},
			output{[][]int{}},
		},
	}

	for _, tc := range tcs {
		got := combinationSum3(tc.input.n, tc.input.k)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
