package two

import (
	"fmt"
	"strconv"
	"strings"
)

func getGameId(game string) int {
	firstColonIdx := strings.Index(game, ":")
	gameIdString := game[5:firstColonIdx]
	gameId, err := strconv.Atoi(gameIdString)
	if err != nil {
		fmt.Printf("Could not find valid game ID integer")
		return 0
	}
	return gameId
}

var maxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func extractSets(game string) []string {
	firstColonIdx := strings.Index(game, ":")
	setsString := game[firstColonIdx+2:]
	return strings.Split(setsString, "; ")
}

func extractBlocks(set string) []string {
	return strings.Split(set, ", ")
}

func extractCountAndColor(block string) (int, string) {
	split := strings.Split(block, " ")
	count, err := strconv.Atoi(split[0])
	if err != nil {
		fmt.Printf("Could not find count of blocks")
		count = 0
	}
	return count, split[1]
}

func isGameValid(game string) bool {
	sets := extractSets(game)
	for _, set := range sets {
		blocks := extractBlocks(set)
		for _, block := range blocks {
			count, color := extractCountAndColor(block)
			if count > maxCubes[color] {
				return false
			}
		}
	}
	return true
}

func RunPartOne(input string) int {
	games := strings.Split(input, "\n")
	var sum int
	for _, game := range games {
		game = strings.TrimSpace(game)
		gameId := getGameId(game)
		if isGameValid(game) {
			sum += gameId
		}
	}
	return sum
}

func CalculatePower(game string) int {
	sets := extractSets(game)
	var maxColorBlocks = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, set := range sets {
		blocks := extractBlocks(set)
		for _, block := range blocks {
			count, color := extractCountAndColor(block)
			if maxColorBlocks[color] < count {
				maxColorBlocks[color] = count
			}
		}
	}
	var power int = 1
	for _, maxCount := range maxColorBlocks {
		power *= maxCount
	}
	return power
}

func RunPartTwo(input string) int {
	games := strings.Split(input, "\n")
	var sum int
	for _, game := range games {
		power := CalculatePower(game)
		sum += power
	}
	return sum
}
