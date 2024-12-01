package days

import (
	"fmt"
	"jayseejay/advent-of-code/years/2023/days/1"
	"jayseejay/advent-of-code/years/2023/days/2"
	"jayseejay/advent-of-code/years/2023/days/3"
	"jayseejay/advent-of-code/years/2023/days/4"
	"jayseejay/advent-of-code/years/2023/days/5"
	"os"
)

func openFile(day int) string {
	var path string = fmt.Sprintf("./years/2023/days/%d/input.txt", day)
	b, err := os.ReadFile(path)
	if err != nil {
		panic("Failed to open file")
	}
	return string(b)
}

func printResult(day int, part int, result int) {
	fmt.Printf("Day %d, Part %d\nResult: %d\n", day, part, result)
}

func One(day int, part int) {
	input := openFile(day)

	if part == 1 {
		sum := one.RunPartOne(input)
		printResult(day, part, sum)
	} else if part == 2 {
		sum := one.RunPartTwo(input)
		printResult(day, part, sum)
	}
}

func Two(day int, part int) {
	input := openFile(day)

	if part == 1 {
		sum := two.RunPartOne(input)
		printResult(day, part, sum)
	} else if part == 2 {
		sum := two.RunPartTwo(input)
		printResult(day, part, sum)
	}
}

func Three(day int, part int) {
	input := openFile(day)

	if part == 1 {
		sum := three.RunPartOne(input)
		printResult(day, part, sum)
	} else if part == 2 {
		sum := three.RunPartTwo(input)
		printResult(day, part, sum)
	}
}

func Four(day int, part int) {
	input := openFile(day)

	if part == 1 {
		sum := four.RunPartOne(input)
		printResult(day, part, sum)
	} else if part == 2 {
		sum := four.RunPartTwo(input)
		printResult(day, part, sum)
	}
}

func Five(day int, part int) {
	input := openFile(day)

	if part == 1 {
		result := five.RunPartOne(input)
		printResult(day, part, result)
	}
	if part == 2 {
		result := five.RunPartTwo(input)
		printResult(day, part, result)
	}
}
