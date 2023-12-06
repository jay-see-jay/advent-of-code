package days

import (
	"fmt"
	"jayseejay/advent-of-code-23/days/one"
	"jayseejay/advent-of-code-23/days/two"
	"os"
)

func One(part int) {
	b, err := os.ReadFile("./days/one/input.txt")
	if err != nil {
		panic("Failed to open file")
	}
	if part == 1 {
		sum := one.RunPartOne(string(b))
		fmt.Printf("Day 1, Part 1\nSum: %d\n", sum)
		return
	}

	if part == 2 {
		sum := one.RunPartTwo(string(b))
		fmt.Printf("Day 1, Part 2\nSum: %d\n", sum)
	}
}

func Two(part int) {
	b, err := os.ReadFile("./days/two/input.txt")
	if err != nil {
		panic("Failed to open file")
	}
	if part == 1 {
		sum := two.RunPartOne(string(b))
		fmt.Printf("Day 2, Part 1\nSum: %d\n", sum)
	}

	if part == 2 {
		sum := two.RunPartTwo(string(b))
		fmt.Printf("Day 2, Part 2\nSum: %d\n", sum)
	}
}
