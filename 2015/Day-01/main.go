package main

import "io/ioutil"

func main() {
	data, _ := ioutil.ReadFile("input_01")
	floor := findFloor(string(data))
	println(floor)
}

func findFloor(instructions string) int {
	floor := 0
	for _, character := range instructions {
		if character == '(' {
			floor++
		} else if character == ')' {
			floor--
		}
	}

	return floor
}
