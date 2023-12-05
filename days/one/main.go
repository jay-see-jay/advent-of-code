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

func RunPartOne(input string) int {
	lines := strings.Split(input, "\n")
	var sum int
	for _, line := range lines {
		sum += ExtractDigits(line)
	}
	return sum
}

var numberWords = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func ExtractDigits2(line string) int {

	var firstDigit string
	var lastDigit string

	var firstDigitIdx int = len(line)
	var lastDigitIdx int = -1

	for word, number := range numberWords {
		firstIdx := strings.Index(line, word)
		if firstIdx >= 0 && firstIdx < firstDigitIdx {
			firstDigit = number
			firstDigitIdx = firstIdx
		}

		lastIdx := strings.LastIndex(line, word)
		if lastIdx >= 0 && lastIdx > lastDigitIdx {
			lastDigit = number
			lastDigitIdx = lastIdx
		}
	}

	characters := strings.Split(line, "")

	for i, char := range characters {
		_, err := strconv.Atoi(char)
		if err != nil {
			continue
		}

		if i < firstDigitIdx {
			firstDigit = char
			firstDigitIdx = i
		}
		if i > lastDigitIdx {
			lastDigit = char
			lastDigitIdx = i
		}
	}

	digits, err := strconv.Atoi(firstDigit + lastDigit)
	if err != nil {
		return 0
	}

	return digits
}

func RunPartTwo(input string) int {
	lines := strings.Split(input, "\n")
	var sum int
	for _, line := range lines {
		sum += ExtractDigits2(line)
	}
	return sum
}
