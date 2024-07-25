package monotoneincreasingdigits

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	n int
}

type output struct {
	expected int
}

func TestMonotoneIncreasingDigits(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{10},
			output{9},
		},
		testcase{
			input{1234},
			output{1234},
		},
		testcase{
			input{332},
			output{299},
		},
	}

	for _, tc := range tcs {
		got := monotoneIncreasingDigits(tc.input.n)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
