package besttimetobuyandsellstockiv

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	k      int
	prices []int
}

type output struct {
	expected int
}

func TestMaxProfit(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{2, []int{2, 4, 1}},
			output{2},
		},
		testcase{
			input{2, []int{3, 2, 6, 5, 0, 3}},
			output{7},
		},
	}

	for _, tc := range tcs {
		got := maxProfit(tc.input.k, tc.input.prices)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
