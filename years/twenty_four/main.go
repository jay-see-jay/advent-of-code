package twenty_twenty_four

import (
	"github.com/jayseejay/advent-of-code/years"
	"github.com/jayseejay/advent-of-code/years/twenty_four/one"
)

func One(day int, part int) {
	sum := one.RunPartOne("")
	years.PrintResult(day, part, sum)
}
