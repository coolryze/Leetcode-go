package queuereconstructionbyheight

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	people [][]int
}

type output struct {
	expected [][]int
}

func TestReconstructQueue(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}},
			output{[][]int{{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1}}},
		},
		testcase{
			input{[][]int{{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4}}},
			output{[][]int{{4, 0}, {5, 0}, {2, 2}, {3, 2}, {1, 4}, {6, 0}}},
		},
	}

	for _, tc := range tcs {
		got := reconstructQueue(tc.input.people)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
