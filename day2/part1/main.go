package main

import (
	"aoc2023"
	"fmt"
	"strings"
)

func main() {
	timer := aoc2023.Time()
	defer timer()

	input := aoc2023.GetInput()
	answer := solve(input)
	fmt.Println(answer)
}

func solve(lines []string) int {
	sumVal := 0

	for _, l := range lines {
		l = strings.ReplaceAll(l, " ", "")
		parts := strings.Split(l, ":")

		game, _ := strings.CutPrefix(parts[0], "Game")
		gameId := aoc2023.Atoi(game)
		if isValidGame(parts[1]) {
			sumVal += gameId
		}
	}

	return sumVal
}

func isValidGame(game string) bool {
	for _, round := range strings.Split(game, ";") {
		isValid := isValidRound(round)
		if !isValid {
			return false
		}
	}

	return true
}

func isValidRound(round string) bool {
	for _, color := range strings.Split(round, ",") {
		if strings.Contains(color, "red") && aoc2023.Atoi(color[:len(color)-3]) > red {
			return false
		}
		if strings.Contains(color, "blue") && aoc2023.Atoi(color[:len(color)-4]) > blue {
			return false
		}
		if strings.Contains(color, "green") && aoc2023.Atoi(color[:len(color)-5]) > green {
			return false
		}
	}

	return true
}

var red int = 12
var green int = 13
var blue int = 14
