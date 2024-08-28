package nextgreaterelementi

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
	expected []int
}

func TestNextGreaterElement(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{4, 1, 2}, []int{1, 3, 4, 2}},
			output{[]int{-1, 3, -1}},
		},
		testcase{
			input{[]int{2, 4}, []int{1, 2, 3, 4}},
			output{[]int{3, -1}},
		},
	}

	for _, tc := range tcs {
		got := nextGreaterElement(tc.input.nums1, tc.input.nums2)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
