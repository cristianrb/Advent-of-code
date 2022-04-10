package main

import "io/ioutil"

func main() {
	moves, _ := ioutil.ReadFile("inputs/input_01")
	houses := countHouses(string(moves))
	println(houses)
}

type Coords struct {
	x int
	y int
}

func countHouses(moves string) int {
	houses := make(map[Coords]bool)

	actualCoords := Coords{0, 0}
	houses[actualCoords] = true

	for _, move := range moves {
		switch move {
		case '^':
			actualCoords.y++
		case '>':
			actualCoords.x++
		case '<':
			actualCoords.x--
		case 'v':
			actualCoords.y--
		}

		houses[actualCoords] = true
	}

	return len(houses)
}
