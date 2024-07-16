package jumpgame

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

func TestCanJump(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{2, 3, 1, 1, 4}},
			output{true},
		},
		testcase{
			input{[]int{3, 2, 1, 0, 4}},
			output{false},
		},
	}

	for _, tc := range tcs {
		got := canJump(tc.input.nums)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
