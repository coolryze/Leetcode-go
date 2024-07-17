package jumpgameii

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

func TestJump(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{2, 3, 1, 1, 4}},
			output{2},
		},
		testcase{
			input{[]int{2, 3, 0, 1, 4}},
			output{2},
		},
		testcase{
			input{[]int{1, 2, 3}},
			output{2},
		},
		testcase{
			input{[]int{1, 1, 1, 1}},
			output{3},
		},
		testcase{
			input{[]int{1, 2, 1, 1, 1}},
			output{3},
		},
		testcase{
			input{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}},
			output{2},
		},
	}

	for _, tc := range tcs {
		got := jump(tc.input.nums)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
