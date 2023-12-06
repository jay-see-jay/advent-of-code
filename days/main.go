package days

import (
	"fmt"
	"jayseejay/advent-of-code-23/days/1"
	"jayseejay/advent-of-code-23/days/2"
	"jayseejay/advent-of-code-23/days/3"
	"os"
)

func One(part int) {
	b, err := os.ReadFile("./days/1/input.txt")
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
	b, err := os.ReadFile("./days/2/input.txt")
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

func Three(part int) {
	b, err := os.ReadFile("./days/3/input.txt")
	if err != nil {
		panic("Failed to open file")
	}

	if part == 1 {
		sum := three.RunPartOne(string(b))
		fmt.Printf("Day 3, Part 1\nSum: %d\n", sum)
	}

	if part == 2 {
		fmt.Printf("Implement me!")
	}
}
