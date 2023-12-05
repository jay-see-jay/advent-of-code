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

var numbers = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
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

func filterPrefix(prefix string) (ret []string) {
	for _, number := range numbers {
		if strings.HasPrefix(number, prefix) {
			ret = append(ret, number)
		}
	}
	return
}

func filterSuffix(suffix string) (ret []string) {
	for _, number := range numbers {
		if strings.HasSuffix(number, suffix) {
			ret = append(ret, number)
		}
	}
	return
}

func checkCharacter(char string, search string, isPrefix bool) (string, string) {
	var numbers []string
	if isPrefix {
		search = search + char
		numbers = filterPrefix(search)
		if len(numbers) == 0 {
			search = char
			numbers = filterPrefix(search)
		}
	} else {
		search = char + search
		numbers = filterSuffix(search)
		if len(numbers) == 0 {
			search = char
			numbers = filterSuffix(search)
		}
	}

	if len(numbers) > 1 {
		return "", search
	}

	if len(numbers) == 1 {
		if numbers[0] == search {
			return numbers[0], ""
		} else {
			return "", search
		}
	}

	return "", ""
}

func ExtractDigits2(line string) int {
	characters := strings.Split(line, "")
	j := len(line) - 1
	var firstDigit string
	var lastDigit string
	var prefix string
	var suffix string
	for i := 0; i < len(characters); i++ {
		_, errI := strconv.Atoi(characters[i])
		if firstDigit == "" {
			if errI == nil {
				firstDigit = characters[i]
			} else {
				numberWord, newPrefix := checkCharacter(characters[i], prefix, true)
				if numberWord != "" {
					firstDigit = numberWords[numberWord]
				}
				prefix = newPrefix
			}
		}

		_, errJ := strconv.Atoi(characters[j])
		if lastDigit == "" {
			if errJ == nil {
				lastDigit = characters[j]
			} else {
				numberWord, newSuffix := checkCharacter(characters[j], suffix, false)
				if numberWord != "" {
					lastDigit = numberWords[numberWord]
				}
				suffix = newSuffix
			}
		}

		j--
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
