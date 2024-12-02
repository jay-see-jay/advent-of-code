package one_test

import (
	"strconv"
	"testing"

	"github.com/jayseejay/advent-of-code/years/twenty_four/one"
)

func TestRunPartOne(t *testing.T) {
	var example = `3   4
4   3
2   5
1   3
3   9
3   3`
	got := one.RunPartOne(example)
	if got != 11 {
		t.Errorf("RunPartOne(3 4, 4 3...) = %d, want 11", got)
	}
}

func TestMeasureDistance(t *testing.T) {
	var tests = []struct {
		name   string
		first  int
		second int
		want   int
	}{
		{"3 4 should be 1", 3, 4, 1},
		{"4 3 should be 1", 4, 3, 1},
		{"3 3 should be 0", 3, 3, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := one.MeasureDistance(test.first, test.second)
			if got != test.want {
				t.Errorf("MeasureDistance(%d, %d) = %s, want %s",
					test.first,
					test.second,
					strconv.Itoa(got),
					strconv.Itoa(test.want))
			}
		})
	}
}

func TestRunPartTwo(t *testing.T) {
	var example = `3   4
4   3
2   5
1   3
3   9
3   3`
	got := one.RunPartTwo(example)
	if got != 31 {
		t.Errorf("RunPartTwo(3 4, 4 3...) = %d, want 31", got)
	}
}
