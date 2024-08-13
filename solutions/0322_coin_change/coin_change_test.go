package coinchange

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	coins  []int
	amount int
}

type output struct {
	expected int
}

func TestCoinChange(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 2, 5}, 11},
			output{3},
		},
		testcase{
			input{[]int{2}, 3},
			output{-1},
		},
		testcase{
			input{[]int{1}, 0},
			output{0},
		},
	}

	for _, tc := range tcs {
		got := coinChange(tc.input.coins, tc.input.amount)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
