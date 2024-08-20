package besttimetobuyandsellstockwithtransactionfee

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
	fee    int
}

type output struct {
	expected int
}

func TestMaxProfit(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 3, 2, 8, 4, 9}, 2},
			output{8},
		},
		testcase{
			input{[]int{1, 3, 7, 5, 10, 3}, 3},
			output{6},
		},
	}

	for _, tc := range tcs {
		got := maxProfit(tc.input.prices, tc.input.fee)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
