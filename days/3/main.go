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

type NumberPosition struct {
	number   int
	position scanner.Position
}

type AdjacentNumber struct {
	number int
	set    bool
}

type AdjacentNumbers struct {
	one AdjacentNumber
	two AdjacentNumber
}

func hasTwoAdjacencies(adjacents AdjacentNumbers) bool {
	return adjacents.one.set && adjacents.two.set
}

type MinMax struct {
	min int
	max int
}

type Surround struct {
	left  int
	right int
}

func isAdjacentPosition(
	numPos NumberPosition,
	lineIdx int,
	colIdx int,
) bool {
	surround := Surround{
		left:  colIdx - 1,
		right: colIdx + 1,
	}
	numPosLineIdx := numPos.position.Line - 1
	numPosColStartIdx := numPos.position.Column - 1
	numberLen := len(strconv.Itoa(numPos.number))
	numPosColEndIdx := numPosColStartIdx + numberLen

	isAbove := numPosLineIdx == lineIdx-1
	isBelow := numPosLineIdx == lineIdx+1
	if isAbove || isBelow {
		isInRange := numPosColStartIdx <= surround.right && numPosColStartIdx >= surround.left-numberLen+1
		return isInRange
	}

	if numPosLineIdx == lineIdx {
		startsImmediatelyAfter := numPosColStartIdx == surround.right
		endsImmediatelyBefore := numPosColEndIdx-1 == surround.left
		return startsImmediatelyAfter || endsImmediatelyBefore
	}

	return false
}

func searchLines(
	lineIdx int,
	colIdx int,
	adjacents *AdjacentNumbers,
	numPos *[]NumberPosition,
) {
	for _, num := range *numPos {
		if isAdjacentPosition(num, lineIdx, colIdx) {
			if !adjacents.one.set {
				adjacents.one.number = num.number
				adjacents.one.set = true
				continue
			}
			if !adjacents.two.set {
				adjacents.two.number = num.number
				adjacents.two.set = true
			}
			if adjacents.one.set && adjacents.two.set {
				break
			}
		}
	}
}

func RunPartTwo(input string) int {
	var s scanner.Scanner
	s.Init(strings.NewReader(input))
	s.Mode = scanner.ScanInts
	s.Whitespace = 1<<' ' | 1<<'\t' | 1<<'\n' | 1<<'.'

	numberPositions := make([]NumberPosition, 0)
	asteriskPositions := make([]scanner.Position, 0)

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		if s.TokenText() != "*" {
			num, err := strconv.Atoi(s.TokenText())
			if err == nil {
				var numPosition = NumberPosition{num, s.Position}
				numberPositions = append(numberPositions, numPosition)
			}
			continue
		}
		asteriskPositions = append(asteriskPositions, s.Position)
	}

	adjacents := AdjacentNumbers{
		one: AdjacentNumber{0, false},
		two: AdjacentNumber{0, false},
	}

	gearRatios := make([]int, 0)

	for _, pos := range asteriskPositions {
		lineIdx := pos.Line - 1
		colIdx := pos.Column - 1

		searchLines(lineIdx, colIdx, &adjacents, &numberPositions)
		if hasTwoAdjacencies(adjacents) {
			gearRatio := adjacents.one.number * adjacents.two.number
			gearRatios = append(gearRatios, gearRatio)
		}
		adjacents.one.set = false
		adjacents.two.set = false
	}

	var sum int
	for _, ratio := range gearRatios {
		sum += ratio
	}

	return sum
}
