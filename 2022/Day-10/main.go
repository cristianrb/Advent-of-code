package main

import (
	"cristianrb-aoc-2022/io"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input, _ := io.ReadLines("Day-10/input_01")
	res1 := strengthSignals(input)
	fmt.Printf("Part 1: %d\n", res1)

	drawCRT(input)
}

func strengthSignals(instructions []string) int {
	cycle := 0
	nCycles := 0
	x := 1
	str := 0

	for _, ins := range instructions {
		cycle, nCycles, str = computeSTRCycle(cycle, nCycles, str, x)
		if strings.Contains(ins, "addx") {
			cycle, nCycles, str = computeSTRCycle(cycle, nCycles, str, x)
			inc, _ := strconv.Atoi(strings.Split(ins, "addx ")[1])
			x += inc
		}
	}

	return str
}

func computeSTRCycle(cycle int, nCycles int, str int, x int) (int, int, int) {
	cycle++
	nCycles++
	if nCycles == 40 || cycle == 20 {
		str += cycle * x
		nCycles = 0
	}
	return cycle, nCycles, str
}

func drawCRT(instructions []string) {
	nCycles := 0
	x := 1
	line := ""

	for _, ins := range instructions {
		nCycles, line = computeCRTCycle(nCycles, x, line)
		if strings.Contains(ins, "addx") {
			nCycles, line = computeCRTCycle(nCycles, x, line)
			inc, _ := strconv.Atoi(strings.Split(ins, "addx ")[1])
			x += inc
		}
	}
}

func computeCRTCycle(nCycles int, x int, line string) (int, string) {
	if nCycles == x || nCycles == x-1 || nCycles == x+1 {
		line += "██"
	} else {
		line += "░░"
	}

	nCycles++
	if nCycles == 40 {
		fmt.Println(line)
		line = ""
		nCycles = 0
	}
	return nCycles, line
}
