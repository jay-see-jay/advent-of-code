package four_test

import (
	"jayseejay/advent-of-code-23/days/4"
	"strings"
	"testing"
)

var example = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

var cardScores = [6]int{8, 2, 2, 1, 0, 0}

func TestRunPartOne(t *testing.T) {
	got := four.RunPartOne(example)

	if got != 13 {
		t.Errorf("RunPartOne(Card 1: 41...) = %d, want 13", got)
	}
}

func TestCalculateCardScore(t *testing.T) {
	cards := strings.Split(example, "\n")

	for i, card := range cards {
		t.Run(card[:6], func(t *testing.T) {
			got := four.CalculateCardScore(card)
			if got != cardScores[i] {
				t.Errorf("CalculateCardScore(%s) = %d, want %d",
					card[:6],
					got,
					cardScores[i])
			}
		})
	}
}

func TestRunPartTwo(t *testing.T) {
	got := four.RunPartTwo(example)
	if got != 30 {
		t.Errorf("RunPartTwo(Card 1: 41...) = %d, want 30", got)
	}
}
