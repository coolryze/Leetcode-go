package permutations

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	nums []int
}

type output struct {
	expected [][]int
}

func TestPermute(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 2, 3}},
			output{[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		},
		testcase{
			input{[]int{0, 1}},
			output{[][]int{{0, 1}, {1, 0}}},
		},
		testcase{
			input{[]int{1}},
			output{[][]int{{1}}},
		},
	}

	for _, tc := range tcs {
		got := permute(tc.input.nums)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
