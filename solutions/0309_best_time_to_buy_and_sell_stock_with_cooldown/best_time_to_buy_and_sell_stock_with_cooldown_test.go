package besttimetobuyandsellstockwithcooldown

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	prices []int
}

type output struct {
	expected int
}

func TestMaxProfit(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 2, 3, 0, 2}},
			output{3},
		},
		testcase{
			input{[]int{1}},
			output{0},
		},
	}

	for _, tc := range tcs {
		got := maxProfit(tc.input.prices)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
