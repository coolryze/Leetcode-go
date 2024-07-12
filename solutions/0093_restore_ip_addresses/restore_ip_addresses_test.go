package restoreipaddresses

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
	expected []string
}

func TestRestoreIpAddresses(t *testing.T) {
	tcs := []testcase{
		testcase{
			input{"25525511135"},
			output{[]string{"255.255.11.135", "255.255.111.35"}},
		},
		testcase{
			input{"0000"},
			output{[]string{"0.0.0.0"}},
		},
		testcase{
			input{"101023"},
			output{[]string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}},
		},
	}

	for _, tc := range tcs {
		got := restoreIpAddresses(tc.input.s)
		if !reflect.DeepEqual(got, tc.output.expected) {
			t.Errorf("test case failed: input=%v, expected=%v, got=%v", tc.input, tc.output.expected, got)
		}
	}
}
