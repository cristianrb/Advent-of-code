package main

import (
	"cristianrb-aoc-2022/io"
	"fmt"
	"strings"
)

func main() {
	data, _ := io.ReadLines("Day-03/input_01")
	priority := rucksackPriorities(data)
	priorityBadges := elfBadges(data)

	fmt.Printf("Part 1: %d\n", priority)
	fmt.Printf("Part 2: %d\n", priorityBadges)
}

const items = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func rucksackPriorities(rucksacks []string) int {
	totalPriority := 0
	for _, rucksack := range rucksacks {
		comp1 := rucksack[:len(rucksack)/2]
		comp2 := rucksack[len(rucksack)/2:]

		for _, c := range comp1 {
			if strings.Contains(comp2, string(c)) {
				totalPriority += strings.IndexRune(items, c) + 1
				break
			}
		}
	}

	return totalPriority

}

func elfBadges(rucksacks []string) int {
	totalPriority := 0
	for i := 0; i < len(rucksacks); i += 3 {
		for _, c := range rucksacks[i] {
			if strings.Contains(rucksacks[i+1], string(c)) && strings.Contains(rucksacks[i+2], string(c)) {
				totalPriority += strings.IndexRune(items, c) + 1
				break
			}
		}
	}

	return totalPriority

}
