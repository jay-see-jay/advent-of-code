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

func isGameValid(game string) bool {
	sets := extractSets(game)
	for _, set := range sets {
		colors := strings.Split(set, ", ")
		for _, color := range colors {
			split := strings.Split(color, " ")
			count, err := strconv.Atoi(split[0])
			if err != nil {
				fmt.Printf("Could not find count of blocks")
				continue
			}
			if count > maxCubes[split[1]] {
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
