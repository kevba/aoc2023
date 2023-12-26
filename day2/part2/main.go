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

		values := gameSet(parts[1])
		sumVal += values[0] * values[1] * values[2]
	}

	return sumVal
}

func gameSet(game string) []int {
	max := []int{0, 0, 0}

	for _, round := range strings.Split(game, ";") {
		numbers := parseNumbers(round)
		for _, i := range []int{0, 1, 2} {

			if numbers[i] > max[i] {
				max[i] = numbers[i]
			}
		}
	}

	return max
}

func parseNumbers(round string) []int {
	numbers := []int{0, 0, 0}

	for _, color := range strings.Split(round, ",") {
		if strings.Contains(color, "red") {
			numbers[0] = aoc2023.Atoi(color[:len(color)-3])
		}
		if strings.Contains(color, "blue") {
			numbers[1] = aoc2023.Atoi(color[:len(color)-4])
		}
		if strings.Contains(color, "green") {
			numbers[2] = aoc2023.Atoi(color[:len(color)-5])
		}
	}

	return numbers
}

var red int = 12
var green int = 13
var blue int = 14
