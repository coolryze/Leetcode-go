package maximumlengthofrepeatedsubarray

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

func TestFindLength(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}},
			output{3},
		},
		testcase{
			input{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}},
			output{5},
		},
	}

	for _, tc := range tcs {
		got := findLength(tc.input.nums1, tc.input.nums2)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
