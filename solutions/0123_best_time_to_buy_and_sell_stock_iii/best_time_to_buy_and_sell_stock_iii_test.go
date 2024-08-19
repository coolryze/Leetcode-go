package besttimetobuyandsellstockiii

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
			input{[]int{3, 3, 5, 0, 0, 3, 1, 4}},
			output{6},
		},
		testcase{
			input{[]int{1, 2, 3, 4, 5}},
			output{4},
		},
		testcase{
			input{[]int{7, 6, 4, 3, 1}},
			output{0},
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
