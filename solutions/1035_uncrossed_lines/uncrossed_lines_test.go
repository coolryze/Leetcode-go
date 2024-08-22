package uncrossedlines

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	nums1 []int
	nums2 []int
}

type output struct {
	expected int
}

func TestMaxUncrossedLines(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 4, 2}, []int{1, 2, 4}},
			output{2},
		},
		testcase{
			input{[]int{2, 5, 1, 2, 5}, []int{10, 5, 2, 1, 5, 2}},
			output{3},
		},
		testcase{
			input{[]int{1, 3, 7, 1, 7, 5}, []int{1, 9, 2, 5, 1}},
			output{2},
		},
	}

	for _, tc := range tcs {
		got := maxUncrossedLines(tc.input.nums1, tc.input.nums2)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
