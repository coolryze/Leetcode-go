package uniquebinarysearchtrees

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

func TestNumTrees(t *testing.T) {
	tcs := []testcase{

		testcase{
			input{3},
			output{5},
		},
		testcase{
			input{1},
			output{1},
		},
	}

	for _, tc := range tcs {
		got := numTrees(tc.input.n)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
