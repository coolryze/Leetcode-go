package nqueens

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
	expected [][]string
}

func TestSolveNQueens(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{4},
			output{[][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}},
		},
		testcase{
			input{1},
			output{[][]string{{"Q"}}},
		},
	}

	for _, tc := range tcs {
		got := solveNQueens(tc.input.n)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
