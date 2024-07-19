package lemonadechange

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	bills []int
}

type output struct {
	expected bool
}

func TestLemonadeChange(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{5, 5, 5, 10, 20}},
			output{true},
		},
		testcase{
			input{[]int{5, 5, 10, 10, 20}},
			output{false},
		},
		testcase{
			input{[]int{5, 5, 5, 5, 10, 5, 10, 10, 10, 20}},
			output{true},
		},
	}

	for _, tc := range tcs {
		got := lemonadeChange(tc.input.bills)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
