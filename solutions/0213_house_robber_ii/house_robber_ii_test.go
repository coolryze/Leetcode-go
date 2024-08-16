package houserobberii

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

func TestRob(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{2, 3, 2}},
			output{3},
		},
		testcase{
			input{[]int{1, 2, 3, 1}},
			output{4},
		},
		testcase{
			input{[]int{1, 2, 3}},
			output{3},
		},
	}

	for _, tc := range tcs {
		got := rob(tc.input.nums)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
