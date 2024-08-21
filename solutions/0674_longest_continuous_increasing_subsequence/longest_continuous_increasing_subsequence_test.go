package longestcontinuousincreasingsubsequence

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
	expected int
}

func TestFindLengthOfLCIS(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 3, 5, 4, 7}},
			output{3},
		},
		testcase{
			input{[]int{2, 2, 2, 2, 2}},
			output{1},
		},
	}

	for _, tc := range tcs {
		got := findLengthOfLCIS(tc.input.nums)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
