package two_test

import (
	"testing"

	"github.com/jayseejay/advent-of-code/years/twenty_four/two"
)

func TestRunPartOne(t *testing.T) {
	var example = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`
	got := two.RunPartOne(example)
	if got != 2 {
		t.Errorf("RunPartOne(7 6 4 2 1...) = %d, want 2", got)
	}
}

func TestIsReportSafe(t *testing.T) {
	var tests = []struct {
		name  string
		input []string
		want  bool
	}{
		{"7 6 4 2 1 should be safe", []string{"7", "6", "4", "2", "1"}, true},
		{"1 2 7 8 9 should be unsafe", []string{"1", "2", "7", "8", "9"}, false},
		{"9 7 6 2 1 should be unsafe", []string{"9", "7", "6", "2", "1"}, false},
		{"1 3 2 4 5 should be unsafe", []string{"1", "3", "2", "4", "5"}, false},
		{"8 6 4 4 1 should be unsafe", []string{"8", "6", "4", "4", "1"}, false},
		{"1 3 6 7 9 should be safe", []string{"1", "3", "6", "7", "9"}, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := two.IsReportSafe(test.input)
			if got != test.want {
				t.Errorf("ExtractDigits(%s) = %t, want %t",
					test.input,
					got,
					test.want)
			}
		})
	}
}
