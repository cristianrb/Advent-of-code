package io

import (
	"bufio"
	"os"
	"strconv"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ReadIntLinesAs2D(path string) [][]int {
	grid := [][]int{}
	lines, _ := ReadLines(path)
	for _, line := range lines {
		row := []int{}
		for _, c := range line {
			n, _ := strconv.Atoi(string(c))
			row = append(row, n)
		}
		grid = append(grid, row)
	}

	return grid
}
