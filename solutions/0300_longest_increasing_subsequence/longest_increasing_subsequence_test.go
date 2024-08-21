package longestincreasingsubsequence

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

func TestLsengthOfLIS(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{10, 9, 2, 5, 3, 7, 101, 18}},
			output{4},
		},
		testcase{
			input{[]int{0, 1, 0, 3, 2, 3}},
			output{4},
		},
		testcase{
			input{[]int{7, 7, 7, 7, 7, 7, 7}},
			output{1},
		},
	}

	for _, tc := range tcs {
		got := lengthOfLIS(tc.input.nums)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
