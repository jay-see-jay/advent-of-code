package main

import (
	"jayseejay/advent-of-code-23/days"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		panic("Please pass just one argument")
	}
	day, err := strconv.Atoi(args[1])
	if err != nil {
		panic("Argument was not a valid integer")
	}

	switch day {
	case 1:
		days.One()
	default:
		panic("Please pass a valid day")
	}
}
