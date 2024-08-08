package partitionequalsubsetsum

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
	expected bool
}

func TestCanPartition(t *testing.T) {
	tcs := []testcase{

		testcase{
			input{[]int{1, 5, 11, 5}},
			output{true},
		},
		testcase{
			input{[]int{1, 2, 3, 5}},
			output{false},
		},
	}

	for _, tc := range tcs {
		got := canPartition(tc.input.nums)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
