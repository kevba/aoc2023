package main

import (
	"aoc2023"
	"testing"
)

func TestSolve(t *testing.T) {
	input := aoc2023.GetTestInput()
	solution := 8

	answer := solve(input)
	if answer != solution {
		t.Errorf("answer %v is not equal to solution %v", answer, solution)
	}
}
