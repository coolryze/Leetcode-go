package assigncookies

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	g []int
	s []int
}

type output struct {
	expected int
}

func TestFindContentChildren(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 2, 3}, []int{1, 1}},
			output{1},
		},
		testcase{
			input{[]int{1, 2}, []int{1, 2, 3}},
			output{2},
		},
	}

	for _, tc := range tcs {
		got := findContentChildren(tc.input.g, tc.input.s)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
