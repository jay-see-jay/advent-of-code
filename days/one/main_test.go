package one_test

import (
	"jayseejay/advent-of-code-23/days/one"
	"strconv"
	"testing"
)

func TestDayOne(t *testing.T) {
	var example = `1abc2
		pqr3stu8vwx
		a1b2c3d4e5f
		treb7uchet`
	got := one.Run(example)
	if got != 142 {
		t.Errorf("Run(1abc2...) = %d, want 142", got)
	}
}

func TestExtractDigits(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"1abc2 should 12", "1abc2", 12},
		{"pqr3stu8vwx should 38", "pqr3stu8vwx", 38},
		{"a1b2c3d4e5f should 12", "a1b2c3d4e5f", 15},
		{"treb7uchet should 77", "treb7uchet", 77},
		{"abc should 0", "abx", 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := one.ExtractDigits(test.input)
			if got != test.want {
				t.Errorf("ExtractDigits(%s) = %s, want %s",
					test.input,
					strconv.Itoa(got),
					strconv.Itoa(test.want))
			}
		})
	}
}
