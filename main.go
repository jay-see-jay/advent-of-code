package main

import (
	"flag"
	"jayseejay/advent-of-code/years/2023/days"
)

func twenty23(day int, part int) {
	switch day {
	case 1:
		days.One(day, part)
	case 2:
		days.Two(day, part)
	case 3:
		days.Three(day, part)
	case 4:
		days.Four(day, part)
	case 5:
		days.Five(day, part)
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
		panic("No implementations for 2024")
	case 2023:
		twenty23(*day, *part)
	default:
		panic("Please pass a valid year")
	}
}
