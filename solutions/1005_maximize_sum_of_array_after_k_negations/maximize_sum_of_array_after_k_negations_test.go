package maximizesumofarrayafterknegations

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
	k    int
}

type output struct {
	expected int
}

func TestLargestSumAfterKNegations(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{4, 2, 3}, 1},
			output{5},
		},
		testcase{
			input{[]int{3, -1, 0, 2}, 3},
			output{6},
		},
		testcase{
			input{[]int{2, -3, -1, 5, -4}, 2},
			output{13},
		},
	}

	for _, tc := range tcs {
		got := largestSumAfterKNegations(tc.input.nums, tc.input.k)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
