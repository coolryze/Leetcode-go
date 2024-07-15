package wigglesubsequence

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

func TestWiggleMaxLength(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 7, 4, 9, 2, 5}},
			output{6},
		},
		testcase{
			input{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}},
			output{7},
		},
		testcase{
			input{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			output{2},
		},
	}

	for _, tc := range tcs {
		got := wiggleMaxLength(tc.input.nums)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
