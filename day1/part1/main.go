package main

import (
	"aoc2023"
	"fmt"
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
		first := rune(-1)
		last := rune(0)

		for _, r := range l {
			if !unicode.IsNumber((r)) {
				continue
			}

			if first == -1 {
				first = r
				last = r
			} else {
				last = r
			}
		}

		n := aoc2023.Atoi((string(first) + string(last)))
		sumVal += n
	}

	return sumVal
}
