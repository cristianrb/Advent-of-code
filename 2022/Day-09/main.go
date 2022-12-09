package main

import (
	"cristianrb-aoc-2022/io"
	"cristianrb-aoc-2022/utils"
	"fmt"
	"image"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, _ := io.ReadLines("Day-09/input_01")
	count := countTailPositions(input, 2)
	fmt.Printf("Part 1: %d\n", count)

	count2 := countTailPositions(input, 10)
	fmt.Printf("Part 2: %d\n", count2)
}

func countTailPositions(input []string, ropeLen int) int {
	rope := make([]image.Point, ropeLen)
	basePoints := map[string]image.Point{"U": {0, -1}, "R": {1, 0}, "D": {0, 1}, "L": {-1, 0}}
	visitedPoints := map[image.Point]bool{}

	for _, instruction := range input {
		insSplit := strings.Split(instruction, " ")
		cmd := insSplit[0]
		steps, _ := strconv.Atoi(insSplit[1])
		for i := 0; i < steps; i++ {
			rope[0] = rope[0].Add(basePoints[cmd])

			for i := 1; i < len(rope); i++ {
				distance := rope[i-1].Sub(rope[i])
				if componentsGreaterThanZero(distance) {
					rope[i] = rope[i].Add(image.Point{X: utils.Sgn(distance.X), Y: utils.Sgn(distance.Y)})
				}
			}

			visitedPoints[rope[len(rope)-1]] = true
		}
	}

	return len(visitedPoints)
}

func componentsGreaterThanZero(d image.Point) bool {
	return math.Abs(float64(d.X)) > 1 || math.Abs(float64(d.Y)) > 1
}
