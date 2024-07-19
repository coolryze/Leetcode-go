package candy

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	ratings []int
}

type output struct {
	expected int
}

func TestCandy(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 0, 2}},
			output{5},
		},
		testcase{
			input{[]int{1, 2, 2}},
			output{4},
		},
	}

	for _, tc := range tcs {
		got := candy(tc.input.ratings)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
