package main

import (
	"cristianrb-aoc-2022/io"
	"fmt"
)

func main() {
	data, _ := io.ReadLines("Day-02/input_01")

	outcomesPart1 := map[string]int{
		"A X": 3 + 1,
		"A Y": 6 + 2,
		"A Z": 3 + 0,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 1 + 6,
		"C Y": 2 + 0,
		"C Z": 3 + 3,
	}

	outcomesPart2 := map[string]int{
		"A X": 3 + 0,
		"A Y": 1 + 3,
		"A Z": 2 + 6,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 6 + 3,
		"C X": 2 + 0,
		"C Y": 3 + 3,
		"C Z": 6 + 1,
	}

	score1 := calcTotalScore(data, outcomesPart1)
	score2 := calcTotalScore(data, outcomesPart2)

	fmt.Printf("Part 1: %d\n", score1)
	fmt.Printf("Part 2: %d\n", score2)
}

func calcTotalScore(games []string, outcomes map[string]int) int {
	score := 0
	for _, game := range games {
		score += outcomes[game]
	}

	return score
}
