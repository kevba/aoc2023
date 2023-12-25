package main

import (
	"aoc2023"
	"fmt"
	"math"
	"strings"
	"unicode"
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
		numberAtIndex := map[int]int{}

		for i, number := range written {
			index := strings.Index(l, number)
			if index != -1 {
				numberAtIndex[index] = i + 1
			}
			lastIndex := strings.LastIndex(l, number)
			if lastIndex != -1 {
				numberAtIndex[lastIndex] = i + 1
			}
		}

		for i, r := range l {
			if !unicode.IsNumber((r)) {
				continue
			}

			numberAtIndex[i] = aoc2023.Atoi((string(r)))
		}

		min := math.MaxInt
		max := 0

		for k := range numberAtIndex {
			if k < min {
				min = k
			}
			if k > max {
				max = k
			}
		}

		if max == 0 {
			max = min
		}

		sumVal += numberAtIndex[min] * 10
		sumVal += numberAtIndex[max]
	}

	return sumVal
}

var written = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
