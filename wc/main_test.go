package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_count(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output [4]int64
	}{
		{
			name: "should return correct counts with valid input",
			input: `Hello,
World!
`,
			output: [4]int64{2, 2, 14, 14},
		},
		{
			name:   "should return 0 for all counts with empty input",
			input:  "",
			output: [4]int64{0, 0, 0, 0},
		},
		{
			name: "should return correct counts with new line only",
			input: `
`,
			output: [4]int64{1, 0, 1, 1},
		},
		{
			name: "should return correct counts with multiple newlines",
			input: `

`,
			output: [4]int64{2, 0, 2, 2},
		},
		{
			name:   "should return correct counts with single line",
			input:  "Hello, World",
			output: [4]int64{1, 2, 12, 12},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lc, wc, bc, chc := count(strings.NewReader(tc.input))

			if got := [4]int64{lc, wc, bc, chc}; !reflect.DeepEqual(got, tc.output) {
				t.Errorf("wc() = %v, want %v", got, tc.output)
			}
		})
	}
}
