package one_test

import (
	"jayseejay/advent-of-code-23/days/one"
	"strconv"
	"testing"
)

func TestRunPartOne(t *testing.T) {
	var example = `1abc2
		pqr3stu8vwx
		a1b2c3d4e5f
		treb7uchet`
	got := one.RunPartOne(example)
	if got != 142 {
		t.Errorf("RunPartOne(1abc2...) = %d, want 142", got)
	}
}

func TestExtractDigits(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"1abc2 should be 12", "1abc2", 12},
		{"pqr3stu8vwx should be 38", "pqr3stu8vwx", 38},
		{"a1b2c3d4e5f should be 12", "a1b2c3d4e5f", 15},
		{"treb7uchet should be 77", "treb7uchet", 77},
		{"abc should be 0", "abx", 0},
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

func TestRunPartTwo(t *testing.T) {
	var example = `two1nine
		eightwothree
		abcone2threexyz
		xtwone3four
		4nineeightseven2
		zoneight234
		7pqrstsixteen`
	got := one.RunPartTwo(example)
	if got != 281 {
		t.Errorf("RunPartOne(1abc2...) = %d, want 281", got)
	}
}

func TestExtractDigits2(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"two1nine should be 29", "two1nine", 29},
		{"eightwothree should be 83", "eightwothree", 83},
		{"abcone2threexyz should be 13", "abcone2threexyz", 13},
		{"xtwone3four should be 24", "xtwone3four", 24},
		{"4nineeightseven2 should be 42", "4nineeightseven2", 42},
		{"zoneight234 should be 14", "zoneight234", 14},
		{"7pqrstsixteen should be 76", "7pqrstsixteen", 76},
		{"abc should be 0", "abx", 0},
		{"oneightwonethree should be 13", "oneightwonethree", 13},
		{"ssevenine eightwoo should be 72", "sevenine eightwo", 72},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := one.ExtractDigits2(test.input)
			if got != test.want {
				t.Errorf("ExtractDigits(%s) = %s, want %s",
					test.input,
					strconv.Itoa(got),
					strconv.Itoa(test.want))
			}
		})
	}
}
