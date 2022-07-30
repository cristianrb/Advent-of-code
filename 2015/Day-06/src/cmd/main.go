package main

import (
	"cristianrb2015/io"
	"strconv"
	"strings"
)

func main() {
	lines, _ := io.ReadLines("Day-06/inputs/input_01")
	println(doOnLights(lines))
}

func doOnLights(input []string) (int, int) {
	lightsMap := map[string]bool{}
	brightnessMap := map[string]int{}
	for _, line := range input {
		action := "turn on"
		if strings.Contains(line, "turn off") {
			action = "turn off"
		} else if strings.Contains(line, "toggle") {
			action = "toggle"
		}
		lights := strings.Split(line, action)[1]
		coords := strings.Split(lights, "through")
		startCoords := strings.Split(coords[0], ",")
		endCoords := strings.Split(coords[1], ",")

		x0, _ := strconv.Atoi(strings.TrimSpace(startCoords[0]))
		y0, _ := strconv.Atoi(strings.TrimSpace(startCoords[1]))
		x1, _ := strconv.Atoi(strings.TrimSpace(endCoords[0]))
		y1, _ := strconv.Atoi(strings.TrimSpace(endCoords[1]))

		for i := x0; i <= x1; i++ {
			for j := y0; j <= y1; j++ {
				coord := strconv.Itoa(i) + "-" + strconv.Itoa(j)
				if action == "turn on" {
					lightsMap[coord] = true
					brightnessMap[coord]++
				} else if action == "turn off" {
					delete(lightsMap, coord)
					if brightnessMap[coord] > 0 {
						brightnessMap[coord]--
					}
				} else if action == "toggle" {
					brightnessMap[coord] += 2
					if !lightsMap[coord] {
						lightsMap[coord] = true
					} else {
						delete(lightsMap, coord)
					}
				}
			}
		}
	}

	brightnessLevel := 0
	for _, v := range brightnessMap {
		brightnessLevel += v
	}

	return len(lightsMap), brightnessLevel
}
