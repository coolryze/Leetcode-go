package coinchangeii

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	amount int
	coins  []int
}

type output struct {
	expected int
}

func TestChange(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{5, []int{1, 2, 5}},
			output{4},
		},
		testcase{
			input{3, []int{2}},
			output{0},
		},
		testcase{
			input{10, []int{10}},
			output{1},
		},
	}

	for _, tc := range tcs {
		got := change(tc.input.amount, tc.input.coins)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
