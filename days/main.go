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
		fmt.Printf("Sum (Part 1): %d\n", sum)
		return
	}

	if part == 2 {
		sum := one.RunPartTwo(string(b))
		fmt.Printf("Sum (Part 2): %d\n", sum)
	}
}

func Two(part int) {
	if part == 1 {
		two.RunPartOne()
	}
}
