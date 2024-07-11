package combinationsum

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

func TestCombinationSum(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{2, 3, 6, 7}, 7},
			output{[][]int{{2, 2, 3}, {7}}},
		},
		testcase{
			input{[]int{2, 3, 5}, 8},
			output{[][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		},
		testcase{
			input{[]int{2}, 1},
			output{[][]int{}},
		},
		testcase{
			input{[]int{8, 7, 4, 3}, 11},
			output{[][]int{{3, 4, 4}, {3, 8}, {4, 7}}},
		},
	}

	for _, tc := range tcs {
		got := combinationSum(tc.input.candidates, tc.input.target)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
