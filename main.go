package main

import (
	"flag"

	twenty_three "github.com/jayseejay/advent-of-code/years/2023/days"
	twenty_four "github.com/jayseejay/advent-of-code/years/twenty_four"
)

func twenty23(day int, part int) {
	switch day {
	case 1:
		twenty_three.One(day, part)
	case 2:
		twenty_three.Two(day, part)
	case 3:
		twenty_three.Three(day, part)
	case 4:
		twenty_three.Four(day, part)
	case 5:
		twenty_three.Five(day, part)
	default:
		panic("Please pass a valid day")
	}
}

func twenty24(day int, part int) {
	switch day {
	case 1:
		twenty_four.One(day, part)
	default:
		panic("Please pass a valid day")
	}
}

func main() {
	year := flag.Int("year", 2024, "The year of the Advent of Code challenge")
	day := flag.Int("day", 1, "The day of the challenge")
	part := flag.Int("part", 1, "Part 1 or 2 of the day's challenge?")

	flag.Parse()

	switch *year {
	case 2024:
		twenty24(*day, *part)
	case 2023:
		twenty23(*day, *part)
	default:
		panic("Please pass a valid year")
	}
}
