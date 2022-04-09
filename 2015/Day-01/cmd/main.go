package main

import "io/ioutil"

func main() {
	data, _ := ioutil.ReadFile("inputs/input_01")
	floor, _ := findFloor(string(data), false)
	println(floor)

	data2, _ := ioutil.ReadFile("inputs/input_02")
	floor, basementChar := findFloor(string(data2), true)
	println(floor)
	println(basementChar)
}

func findFloor(instructions string, secondPart bool) (int, int) {
	floor := 0
	basementLevel := -1
	for index, character := range instructions {
		if character == '(' {
			floor++
		} else if character == ')' {
			floor--
		}

		if floor == basementLevel && secondPart {
			return floor, index + 1
		}

	}

	return floor, -1
}
