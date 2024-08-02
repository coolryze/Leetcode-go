package fibonaccinumber

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

func TestFib(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{2},
			output{1},
		},
		testcase{
			input{3},
			output{2},
		},
		testcase{
			input{4},
			output{3},
		},
	}

	for _, tc := range tcs {
		got := fib(tc.input.n)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
