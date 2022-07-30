package main

import "io/ioutil"

func main() {
	moves, _ := ioutil.ReadFile("Day-03/inputs/input_01")
	houses := countHouses(string(moves))
	println(len(houses))

	moves2, _ := ioutil.ReadFile("Day-03/inputs/input_02")
	santaMoves := ""
	roboSantaMoves := ""
	for i, move := range string(moves2) {
		if i%2 == 0 {
			santaMoves += string(move)
		} else {
			roboSantaMoves += string(move)
		}
	}

	santaHouses := countHouses(santaMoves)
	roboSantaHouses := countHouses(roboSantaMoves)

	fullHouses := make(map[Coords]bool)
	for k, v := range santaHouses {
		fullHouses[k] = v
	}
	for k, v := range roboSantaHouses {
		fullHouses[k] = v
	}
	println(len(fullHouses))
}

type Coords struct {
	x int
	y int
}

func countHouses(moves string) map[Coords]bool {
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

	return houses
}
