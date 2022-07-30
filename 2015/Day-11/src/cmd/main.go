package main

import (
	"fmt"
	"regexp"
)

func main() {
	input := "hxbxxyzz"
	firstSolution := nextPassword(input)
	secondSolution := nextPassword(firstSolution)
	println(firstSolution)
	println(secondSolution)
}

func nextPassword(input string) string {
	nextPwd := incrementString(input)
	notValid := isValidPassword(nextPwd)
	for !notValid {
		nextPwd = incrementString(nextPwd)
		notValid = isValidPassword(nextPwd)
	}
	return nextPwd
}

func incrementChar(character rune) rune {
	if character == 'z' {
		return 'a'
	}
	return character + 1
}

func incrementString(input string) string {
	lastChar := input[len(input)-1]
	newChar := incrementChar(rune(lastChar))
	if newChar == 'a' {
		inputMinusOne := input[0 : len(input)-1]
		newString := fmt.Sprintf("%s%s", incrementString(inputMinusOne), "a")
		return newString
	} else {
		return fmt.Sprintf("%s%s", input[0:len(input)-1], string(newChar))
	}
}

func isValidPassword(password string) bool {
	containsNotValidSymbols, _ := regexp.MatchString("i|o|l", password)
	if containsNotValidSymbols {
		return false
	}

	if !hasTwoDifferentPairs(password) {
		return false
	}
	for idx, _ := range password {
		if idx+1 >= len(password) || idx+2 >= len(password) {
			return false
		}
		if password[idx]+1 == password[idx+1] && password[idx+1]+1 == password[idx+2] {
			return true
		}
	}
	return false

}

func hasTwoDifferentPairs(input string) bool {
	pairs := map[string]bool{}
	for i := 0; i < len(input)-1; i++ {
		if input[i] == input[i+1] {
			_, exist := pairs[string(input[i])]
			if !exist {
				pairs[string(input[i])] = true
			}
		}

		if len(pairs) > 1 {
			return true
		}
	}
	return false
}
