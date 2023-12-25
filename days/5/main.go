package five

import (
	"fmt"
	"strconv"
	"strings"
)

// For each seed, map:
// - seed to soil
// - soil to fertiliser
// - fertiliser to water
// - water to light
// - light to temperature
// - temperature to humidity
// - humidity to location

// For each line of each map:
// - If input number >= second number and < (second number + third number):
// -- second number to (second number + (second number + third number - 1))
// -- maps to...
// -- first number to (first number + (first number + third number - 1))
// - Else:
// -- return input number

// A map should take an input number:
// - check if it falls within any of the ranges
// - if yes apply modifier (second number - first number)

// map of maps:

type StartEnd struct {
	start int
	end   int
}

type MapRange struct {
	source      StartEnd
	destination StartEnd
	modifier    int
}

var maps = make(map[string][]MapRange)

// Go through input string to:
// - extract seed numbers
// -- if start of input, create list of numbers after first colon up to first \n
// - populate `maps`
// -- after seed numbers, two \n denotes a new map, get map key from text immediately after up to work "map:"
// -- each subsequent new line is a map range

var seeds = make([]int, 0)

func isNumber(char rune) bool {
	return char >= 48 && char < 58
}

func startNewWord(char rune) bool {
	return char == 32 || char == 10 || char == 58
}

func isNumberComplete(word string) bool {
	if word == "" || strings.HasPrefix(word, "map") {
		return false
	}
	return true
}

func readInput(input string) {
	newLineCount := 0
	key := ""
	isKeyComplete := false
	number := ""
	for _, char := range input {
		stringChar := string(char)
		if char == 10 {
			newLineCount += 1
		} else {
			newLineCount = 0
		}

		if newLineCount == 2 {
			key = ""
			isKeyComplete = false
		}

		if !isNumber(char) && !isKeyComplete {
			if char == 32 || char == 58 {
				isKeyComplete = true
			} else {
				key += stringChar
			}
		}

		if !isKeyComplete {
			continue
		}

		if !startNewWord(char) {
			number += stringChar
			continue
		}

		if !isNumberComplete(number) {
			number = ""
			continue
		}

		if strings.TrimSpace(key) == "seeds" {
			seed, err := strconv.Atoi(number)
			if err == nil {
				fmt.Printf("Append to seeds: %d\n", seed)
				seeds = append(seeds, seed)
			}
		}
		fmt.Printf("Key: %s | Word: %s\n", strings.TrimSpace(key), strings.TrimSpace(number))
		number = ""
	}
}

func RunPartOne(input string) int {
	fmt.Printf("Seeds: %v\n", seeds)
	readInput(input)
	fmt.Printf("Seeds: %v", seeds)
	return 0
}
