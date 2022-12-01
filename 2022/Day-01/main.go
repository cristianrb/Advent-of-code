package main

import (
	"cristianrb-aoc-2022/io"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	data, _ := io.ReadLines("Day-01/input_01")
	caloriesOrdered := mapCaloriesToIntOrdered(data)

	fmt.Printf("Part 1: %d\n", caloriesOrdered[0])
	fmt.Printf("Part 2: %d\n", caloriesOrdered[0]+caloriesOrdered[1]+caloriesOrdered[2])
}

func mapCaloriesToIntOrdered(lines []string) []int {
	actualCalories := 0
	caloriesList := []int{}
	for _, line := range lines {
		if line == "" {
			caloriesList = append(caloriesList, actualCalories)
			actualCalories = 0
		} else {
			calories, _ := strconv.Atoi(line)
			actualCalories += calories
		}
	}
	caloriesList = append(caloriesList, actualCalories)

	sort.Slice(caloriesList, func(i, j int) bool {
		return caloriesList[i] > caloriesList[j]
	})

	return caloriesList
}
