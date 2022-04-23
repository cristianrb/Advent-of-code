package main

import (
	"cristianrbaoc2015d5/src/io"
)

func main() {
	lines, _ := io.ReadLines("inputs/input_01")
	counter := 0
	for _, line := range lines {
		if isNiceV2(line) {
			counter++
		}
	}
	println(counter)
}

func isNice(input string) bool {
	vowels := []string{"a", "e", "i", "o", "u"}
	invalidSequences := []string{"ab", "cd", "pq", "xy"}
	vowelsCounter := 0
	doubleLetter := false
	var lastCharacter uint8
	for i := 0; i < len(input); i++ {
		actualCharacter := input[i]
		if contains(vowels, string(actualCharacter)) {
			vowelsCounter++
		}
		if i != 0 {
			if lastCharacter == actualCharacter {
				doubleLetter = true
			}
			invalidString := string(lastCharacter) + string(actualCharacter)
			if contains(invalidSequences, invalidString) {
				return false
			}
		}
		lastCharacter = actualCharacter
	}

	return vowelsCounter >= 3 && doubleLetter
}

func isNiceV2(input string) bool {
	repeatLetter := false
	pairTwice := false
	pairs := map[string]int{}

	for i := 0; i < len(input)-1; i++ {
		pair := string(input[i]) + string(input[i+1])
		previousIndex, exist := pairs[pair]
		if exist && previousIndex != i && previousIndex != i-1 {
			pairTwice = true
		}

		if !exist {
			pairs[pair] = i
		}

		if i+2 < len(input) {
			if input[i] == input[i+2] {
				repeatLetter = true
			}
		}

	}

	return repeatLetter && pairTwice
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
