package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func parseCol(in []string, col int) []string {
	stack := []string{}

	for i := 0; i < len(in); i++ {
		if l := in[i][col]; unicode.IsLetter(rune(l)) {
			stack = append([]string{string(l)}, stack...)
		}
	}

	return stack
}

func main() {
	data, _ := os.ReadFile("Day-05/input_01")
	println(data)
	stacks := strings.Split(strings.Split(string(data), "\r\n\r\n")[0], "\r\n")
	instructions := strings.Split(strings.Split(string(data), "\r\n\r\n")[1], "\r\n")

	inputMap := map[int][]string{}
	inputMap2 := map[int][]string{}
	for i, imI := 1, 1; i < len(stacks[0]); i += 4 {
		stack := parseCol(stacks, i)
		stack2 := parseCol(stacks, i) // could just copy it...
		inputMap[imI] = stack
		inputMap2[imI] = stack2
		imI++
	}

	res1 := topCrates(inputMap, instructions, false)
	fmt.Printf("Part 1: %s\n", res1)

	res2 := topCrates(inputMap2, instructions, true)
	fmt.Printf("Part 2: %s\n", res2)
}

func topCrates(input map[int][]string, instructions []string, preserveOrder bool) string {
	for _, command := range instructions {
		strS := strings.Split(command, " from ")
		actionStr := strings.Split(strS[0], "move ")[1]
		action, _ := strconv.Atoi(actionStr)
		stacks := strings.Split(strS[1], " to ")
		fromStack, _ := strconv.Atoi(stacks[0])
		toStack, _ := strconv.Atoi(stacks[1])

		if preserveOrder {
			cratesToAdd := []string{}
			for i := 0; i < action; i++ {
				lastElem := len(input[fromStack]) - 1
				crate := input[fromStack][lastElem]
				cratesToAdd = append(cratesToAdd, crate)
				input[fromStack] = input[fromStack][0:lastElem]
			}

			for i := len(cratesToAdd) - 1; i >= 0; i-- {
				input[toStack] = append(input[toStack], cratesToAdd[i])
			}
		} else {
			for i := 0; i < action; i++ {
				lastElem := len(input[fromStack]) - 1
				crate := input[fromStack][lastElem]
				input[fromStack] = input[fromStack][0:lastElem]
				input[toStack] = append(input[toStack], crate)
			}
		}

	}

	res := ""
	keys := make([]int, 0)
	for k, _ := range input {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		lastElem := len(input[k]) - 1
		res += input[k][lastElem]
	}
	return res
}
