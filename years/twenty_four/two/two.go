package two

import (
	"strconv"
	"strings"
	"text/scanner"
)

type callback func(string)

// func readCallback(token string, sum *int) int {
// 	return 0
// }

func readInput(input string, fn callback) int {
	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	s.Mode = scanner.ScanInts
	s.Whitespace = 1<<' ' | 1<<'\t' | 1<<'\n' | 1<<'.'

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fn(s.TokenText())
	}

	return 0
}

func IsReportSafe(levels []string) bool {
	prev, err := strconv.Atoi(levels[0])
	if err != nil {
		return false
	}

	curr, err := strconv.Atoi(levels[1])
	if err != nil {
		return false
	}
	levels = levels[2:]

	if prev == curr {
		return false
	}

	isAsc := curr > prev

	for _, level := range levels {
		if err != nil {
			return false
		}
		next, nextErr := strconv.Atoi(level)

		var difference int

		if isAsc {
			difference = next - curr
		} else {
			difference = curr - next
		}

		if difference < 1 || difference > 3 {
			return false
		}

		prev = curr
		curr = next
		err = nextErr
	}

	return true
}

func RunPartOne(input string) int {
	reports := strings.Split(input, "\n")

	safeCount := 0

	for _, report := range reports {
		levels := strings.Split(report, " ")
		if IsReportSafe(levels) {
			safeCount++
		}
	}

	return safeCount
}
