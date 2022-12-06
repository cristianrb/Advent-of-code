package main

import (
	"fmt"
	"os"
)

func main() {
	input, _ := os.ReadFile("Day-06/input_01")
	res := firstNUniqueChars(string(input), 4)
	res2 := firstNUniqueChars(string(input), 14)

	fmt.Printf("Part 1: %d\n", res)
	fmt.Printf("Part 2: %d\n", res2)
}

func firstNUniqueChars(input string, n int) int {
	lastFour := input[0:n]
	if uniqueChars(lastFour) {
		return n
	}
	lastIdx := 1
	for i := n; i < len(input); i++ {
		if uniqueChars(input[lastIdx : i+1]) {
			return i + 1
		}
		lastIdx++
	}
	return -1
}

func uniqueChars(str string) bool {
	letters := map[rune]bool{}
	for _, l := range str {
		if letters[l] {
			return false
		}
		letters[l] = true
	}

	return true
}
