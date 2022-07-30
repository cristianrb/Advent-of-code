package main

import (
	"cristianrb2015/io"
	"cristianrb2015/util"
	"strconv"
	"strings"
)

type Direction struct {
	Origin      string
	Destination string
	Distance    int
}

func main() {
	lines, _ := io.ReadLines("Day-09/inputs/input_01")
	cities, directions := generateCitiesAndDirections(lines)
	permutations := generatePermutations(cities)
	distances := computeDistances(permutations, directions)
	// Sol 1
	println(util.MinOf(distances...))
	// Sol 2
	println(util.MaxOf(distances...))
}

func computeDistances(permutations [][]string, directions []Direction) []int {
	var distances []int
	for _, cities := range permutations {
		distance := 0
		origin := cities[0]
		for idx := 1; idx < len(cities); idx++ {
			destination := cities[idx]
			for _, direction := range directions {
				if directionMatches(origin, destination, direction) {
					distance += direction.Distance
				}
			}
			origin = destination
		}
		distances = append(distances, distance)
	}
	return distances
}

func directionMatches(origin string, destination string, direction Direction) bool {
	return (origin == direction.Origin || origin == direction.Destination) && (destination == direction.Origin || destination == direction.Destination)
}

func generatePermutations(cities []string) [][]string {
	var permutations [][]string
	util.Perm(cities, func(cities []string) {
		var permutation []string
		for _, city := range cities {
			permutation = append(permutation, city)
		}
		permutations = append(permutations, permutation)
	})
	return permutations
}

func generateCitiesAndDirections(lines []string) ([]string, []Direction) {
	var directions []Direction
	var cities []string
	for _, direction := range lines {
		directionSplit := strings.Split(direction, " to ")
		origin := directionSplit[0]
		directionSplit = strings.Split(directionSplit[1], " = ")
		destination := directionSplit[0]
		cities = addOriginAndDestination(cities, origin, destination)
		distance, _ := strconv.Atoi(directionSplit[1])
		directions = append(directions, Direction{
			origin,
			destination,
			distance,
		})
	}
	return cities, directions
}

func addOriginAndDestination(cities []string, origin string, destination string) []string {
	if !util.Contains(cities, origin) {
		cities = append(cities, origin)
	}
	if !util.Contains(cities, destination) {
		cities = append(cities, destination)
	}
	return cities
}
