package three_test

import (
	"testing"

	"github.com/jayseejay/advent-of-code/years/twenty_three/three"
)

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

func TestRunPartOne(t *testing.T) {
	got := three.RunPartOne(example)

	if got != 4361 {
		t.Errorf("RunPartOne(467..114...) = %d, want 4361", got)
	}

}

func TestRunPartTwo(t *testing.T) {
	got := three.RunPartTwo(example)
	if got != 467835 {
		t.Errorf("RunPartTwo(467..114...) = %d, want 467835", got)
	}
}
