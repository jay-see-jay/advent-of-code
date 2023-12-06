package three_test

import (
	"jayseejay/advent-of-code-23/days/3"
	"testing"
)

func TestRunPartOne(t *testing.T) {
	var example = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	got := three.RunPartOne(example)

	if got != 4361 {
		t.Errorf("RunPartOne(467..114...) = %d, want 4361", got)
	}

}
