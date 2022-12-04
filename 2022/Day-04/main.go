package main

import (
	"cristianrb-aoc-2022/io"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data, _ := io.ReadLines("Day-04/input_01")
	count := 0
	count2 := 0
	for _, line := range data {
		if isFullyContained(line) {
			count++
		}
		if isContained(line) {
			count2++
		}
	}

	fmt.Printf("Part 1: %d\n", count)
	fmt.Printf("Part 2: %d\n", count2)
}

func isFullyContained(input string) bool {
	pairs := strings.Split(input, ",")
	n1p1, _ := strconv.Atoi(strings.Split(pairs[0], "-")[0])
	n1p2, _ := strconv.Atoi(strings.Split(pairs[0], "-")[1])
	n2p1, _ := strconv.Atoi(strings.Split(pairs[1], "-")[0])
	n2p2, _ := strconv.Atoi(strings.Split(pairs[1], "-")[1])

	return (n1p1 >= n2p1 && n1p2 <= n2p2) || (n2p1 >= n1p1 && n2p2 <= n1p2)
}

func isContained(input string) bool {
	pairs := strings.Split(input, ",")
	n1p1, _ := strconv.Atoi(strings.Split(pairs[0], "-")[0])
	n1p2, _ := strconv.Atoi(strings.Split(pairs[0], "-")[1])
	n2p1, _ := strconv.Atoi(strings.Split(pairs[1], "-")[0])
	n2p2, _ := strconv.Atoi(strings.Split(pairs[1], "-")[1])

	return n1p2 >= n2p1 && n1p1 <= n2p2
}
