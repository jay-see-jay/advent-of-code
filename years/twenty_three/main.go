package twenty_three

import (
	"github.com/jayseejay/advent-of-code/years"
	"github.com/jayseejay/advent-of-code/years/twenty_three/five"
	"github.com/jayseejay/advent-of-code/years/twenty_three/four"
	"github.com/jayseejay/advent-of-code/years/twenty_three/one"
	"github.com/jayseejay/advent-of-code/years/twenty_three/three"
	"github.com/jayseejay/advent-of-code/years/twenty_three/two"
)

func One(day int, part int) {
	input := years.OpenFile(day)

	if part == 1 {
		sum := one.RunPartOne(input)
		years.PrintResult(day, part, sum)
	} else if part == 2 {
		sum := one.RunPartTwo(input)
		years.PrintResult(day, part, sum)
	}
}

func Two(day int, part int) {
	input := years.OpenFile(day)

	if part == 1 {
		sum := two.RunPartOne(input)
		years.PrintResult(day, part, sum)
	} else if part == 2 {
		sum := two.RunPartTwo(input)
		years.PrintResult(day, part, sum)
	}
}

func Three(day int, part int) {
	input := years.OpenFile(day)

	if part == 1 {
		sum := three.RunPartOne(input)
		years.PrintResult(day, part, sum)
	} else if part == 2 {
		sum := three.RunPartTwo(input)
		years.PrintResult(day, part, sum)
	}
}

func Four(day int, part int) {
	input := years.OpenFile(day)

	if part == 1 {
		sum := four.RunPartOne(input)
		years.PrintResult(day, part, sum)
	} else if part == 2 {
		sum := four.RunPartTwo(input)
		years.PrintResult(day, part, sum)
	}
}

func Five(day int, part int) {
	input := years.OpenFile(day)

	if part == 1 {
		result := five.RunPartOne(input)
		years.PrintResult(day, part, result)
	}
	if part == 2 {
		result := five.RunPartTwo(input)
		years.PrintResult(day, part, result)
	}
}
