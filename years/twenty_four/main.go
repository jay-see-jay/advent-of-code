package twenty_twenty_four

import (
	"github.com/jayseejay/advent-of-code/years"
	"github.com/jayseejay/advent-of-code/years/twenty_four/one"
)

func One(day int, part int) {
	input := years.OpenFile(2024, day)
	var sum int
	if part == 1 {
		sum = one.RunPartOne(input)
	} else {
		panic("Part 2 is not yet implemented")
	}
	years.PrintResult(day, part, sum)
}
