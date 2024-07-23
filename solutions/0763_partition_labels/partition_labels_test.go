package partitionlabels

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	s string
}

type output struct {
	expected []int
}

func TestPartitionLabels(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"ababcbacadefegdehijhklij"},
			output{[]int{9, 7, 8}},
		},
		testcase{
			input{"eccbbbbdec"},
			output{[]int{10}},
		},
	}

	for _, tc := range tcs {
		got := partitionLabels(tc.input.s)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
