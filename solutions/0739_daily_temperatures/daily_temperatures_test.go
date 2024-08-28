package dailytemperatures

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	temperatures []int
}

type output struct {
	expected []int
}

func TestDailyTemperatures(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{73, 74, 75, 71, 69, 72, 76, 73}},
			output{[]int{1, 1, 4, 2, 1, 1, 0, 0}},
		},
		testcase{
			input{[]int{30, 40, 50, 60}},
			output{[]int{1, 1, 1, 0}},
		},
		testcase{
			input{[]int{30, 60, 90}},
			output{[]int{1, 1, 0}},
		},
	}

	for _, tc := range tcs {
		got := dailyTemperatures(tc.input.temperatures)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
