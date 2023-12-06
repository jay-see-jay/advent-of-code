package main

import (
	"jayseejay/advent-of-code-23/days"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) < 2 || len(args) > 3 {
		panic("Please pass either one or two arguments")
	}
	day, err := strconv.Atoi(args[1])
	if err != nil {
		panic("Argument was not a valid integer")
	}

	var part int
	if len(args) == 3 {
		num, err := strconv.Atoi(args[2])
		if err != nil || (num < 1 || num > 2) {
			panic("Argument was not a valid integer")
		}
		part = num
	} else {
		part = 1
	}

	switch day {
	case 1:
		days.One(part)
	case 2:
		days.Two(part)
	case 3:
		days.Three(part)
	default:
		panic("Please pass a valid day")
	}
}
