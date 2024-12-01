package two_test

import (
	"testing"

	"github.com/jayseejay/advent-of-code/years/2023/days/two"
)

var example = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func TestRunPartOne(t *testing.T) {
	got := two.RunPartOne(example)

	if got != 8 {
		t.Errorf("RunPartOne(Game 1: 3 blue...) = %d, want 8", got)
	}
}

func TestRunPartTwo(t *testing.T) {
	got := two.RunPartTwo(example)

	if got != 2286 {
		t.Errorf("RunPartOne(Game 1: 3 blue...) = %d, want 2286", got)
	}
}

type Test = struct {
	name  string
	input string
	want  int
}

func TestCalculatePower(t *testing.T) {
	var tests = []Test{
		{"Game 1 should be 48", "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 48},
		{"Game 2 should be 12", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 12},
		{"Game 3 should be 1560", "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 1560},
		{"Game 4 should be 630", "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 630},
		{"Game 5 should be 36", "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 36},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := two.CalculatePower(test.input)
			if got != test.want {
				t.Errorf("CalculatePower(%s) = %d, want %d",
					test.input,
					got,
					test.want)
			}
		})
	}

}
