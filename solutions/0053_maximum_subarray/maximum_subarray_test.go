package maximumsubarray

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

func TestMaxSubArray(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			output{6},
		},
		testcase{
			input{[]int{1}},
			output{1},
		},
		testcase{
			input{[]int{5, 4, -1, 7, 8}},
			output{23},
		},
	}

	for _, tc := range tcs {
		got := maxSubArray(tc.input.nums)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
