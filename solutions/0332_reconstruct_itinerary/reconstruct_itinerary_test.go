package reconstructitinerary

import (
	"reflect"
	"testing"
)

type testcase struct {
	input
	output
}

type input struct {
	tickets [][]string
}

type output struct {
	expected []string
}

func TestFindItinerary(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{[][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}},
			output{[]string{"JFK", "MUC", "LHR", "SFO", "SJC"}},
		},
	}

	for _, tc := range tcs {
		got := findItinerary(tc.input.tickets)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
