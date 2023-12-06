package three

import (
	"strconv"
	"strings"
	"text/scanner"
)

func getJoiningRow(
	startIdx int,
	endIdx int,
	lineIdx int,
	line string,
	lines []string,
	isAbove bool) []string {

	start := max(startIdx-1, 0)
	end := min(endIdx+1, len(line))
	var row string

	if isAbove {
		row = lines[lineIdx-1][start:end]
	} else {
		row = lines[lineIdx+1][start:end]

	}
	return strings.Split(row, "")
}

func checkForNonStop(chars []string) bool {
	for _, char := range chars {
		if char != "." {
			return true
		}
	}
	return false
}

func RunPartOne(input string) int {
	var sum int
	lines := strings.Split(input, "\n")

	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	s.Mode = scanner.ScanInts
	s.Whitespace = 1<<' ' | 1<<'\t' | 1<<'\n' | 1<<'.'

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		num, err := strconv.Atoi(s.TokenText())
		if err != nil {
			continue
		}
		numStartIdx := s.Position.Column - 1
		numEndIdx := numStartIdx + len(s.TokenText())
		lineIdx := s.Position.Line - 1
		line := lines[lineIdx]

		if numStartIdx > 0 && line[numStartIdx-1] != '.' {
			sum += num
			continue
		}

		if numEndIdx < len(line) && line[numEndIdx] != '.' {
			sum += num
			continue
		}

		if lineIdx > 0 {
			characters := getJoiningRow(numStartIdx, numEndIdx, lineIdx, line, lines, true)
			if checkForNonStop(characters) {
				sum += num
			}
		}

		if lineIdx < len(lines)-1 {
			characters := getJoiningRow(numStartIdx, numEndIdx, lineIdx, line, lines, false)
			if checkForNonStop(characters) {
				sum += num
			}
		}
	}

	return sum
}
