package four

import (
	"math"
	"strings"
)

func isNumber(char uint8) bool {
	return char >= '0' && char <= '9'
}

func convertToInt(utf []uint8) (sum int) {
	multiplier := 1
	for i := len(utf) - 1; i >= 0; i-- {
		num := int(utf[i] - 48)
		sum += num * multiplier
		multiplier *= 10
	}
	return
}

func CalculateCardScore(card string) int {
	indexOfColon := strings.Index(card, ":")
	curNumStr := make([]uint8, 0)
	passedVertLine := false
	winningNumbers := make(map[int]struct{})
	matchedNumbers := make([]int, 0)
	for i := indexOfColon; i < len(card); i++ {
		char := card[i]

		if char == '|' {
			passedVertLine = true
			continue
		}

		if isNumber(char) {
			curNumStr = append(curNumStr, char)
		}

		if (char == ' ' || i == len(card)-1) && len(curNumStr) > 0 {
			num := convertToInt(curNumStr)
			if !passedVertLine {
				winningNumbers[num] = struct{}{}
			} else {
				_, ok := winningNumbers[num]
				if ok {
					matchedNumbers = append(matchedNumbers, num)
				}
			}
			curNumStr = make([]uint8, 0)
		}
	}
	sum := 0
	for range matchedNumbers {
		if sum == 0 {
			sum += 1
		} else {
			sum *= 2
		}
	}
	return sum
}

func RunPartOne(input string) (sum int) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		score := CalculateCardScore(line)
		sum += score
	}
	return
}

func RunPartTwo(input string) int {
	lines := strings.Split(input, "\n")
	cards := make(map[int]int)
	for i := range lines {
		cards[i] = 1
	}
	for i, line := range lines {
		score := CalculateCardScore(line)
		if score == 0 {
			continue
		}
		matchedCardsCount := math.Log2(float64(score)) + 1
		for j := 1; matchedCardsCount > 0; j++ {
			cards[i+j] += 1 * cards[i]
			matchedCardsCount--
		}
	}
	sum := 0
	for _, count := range cards {
		sum += count
	}
	return sum
}
