package uniquepathsii

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	obstacleGrid [][]int
}

type output struct {
	expected int
}

func TestUniquePathsWithObstacles(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
			output{2},
		},
		testcase{
			input{[][]int{{0, 1}, {0, 0}}},
			output{1},
		},
		testcase{
			input{[][]int{{1, 0}}},
			output{0},
		},
		testcase{
			input{[][]int{{1}, {0}}},
			output{0},
		},
	}

	for _, tc := range tcs {
		got := uniquePathsWithObstacles(tc.input.obstacleGrid)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
