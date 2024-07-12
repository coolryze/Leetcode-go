package combinationsumii

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	candidates []int
	target     int
}

type output struct {
	expected [][]int
}

func TestCombinationSum2(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{10, 1, 2, 7, 6, 1, 5}, 8},
			output{[][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		},
		testcase{
			input{[]int{2, 5, 2, 1, 2}, 5},
			output{[][]int{{1, 2, 2}, {5}}},
		},
	}

	for _, tc := range tcs {
		got := combinationSum2(tc.input.candidates, tc.input.target)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
