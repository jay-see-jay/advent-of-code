package one

import (
	"strconv"
	"strings"
)

func ExtractDigits(line string) int {
	characters := strings.Split(line, "")
	j := len(line) - 1
	var firstDigit string
	var lastDigit string
	for i := 0; i < len(characters); i++ {
		_, errI := strconv.Atoi(characters[i])
		if firstDigit == "" && errI == nil {
			firstDigit = characters[i]
		}
		_, errJ := strconv.Atoi(characters[j])
		if lastDigit == "" && errJ == nil {
			lastDigit = characters[j]
		}
		j--
	}
	digits, err := strconv.Atoi(firstDigit + lastDigit)
	if err != nil {
		return 0
	}

	return digits
}

func Run(input string) int {
	lines := strings.Split(input, "\n")
	var sum int
	for _, line := range lines {
		sum += ExtractDigits(line)
	}
	return sum
}
