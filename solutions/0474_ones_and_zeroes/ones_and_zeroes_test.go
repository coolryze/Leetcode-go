package onesandzeroes

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	strs []string
	m    int
	n    int
}

type output struct {
	expected int
}

func TestFindMaxForm(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]string{"10", "0001", "111001", "1", "0"}, 5, 3},
			output{4},
		},
		testcase{
			input{[]string{"10", "0", "1"}, 1, 1},
			output{2},
		},
	}

	for _, tc := range tcs {
		got := findMaxForm(tc.input.strs, tc.input.m, tc.input.n)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
