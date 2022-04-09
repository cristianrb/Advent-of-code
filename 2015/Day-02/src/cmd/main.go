package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := readLines("inputs/input_01")
	res := findResFromList(data, findSurface)
	println(res)

	dataSecond, _ := readLines("inputs/input_02")
	resSecond := findResFromList(dataSecond, findRibbon)
	println(resSecond)
}

func findResFromList(data []string, function func(l, w, h int) int) int {
	res := 0
	for _, line := range data {
		lineSplit := strings.Split(line, "x")
		l, _ := strconv.Atoi(lineSplit[0])
		w, _ := strconv.Atoi(lineSplit[1])
		h, _ := strconv.Atoi(lineSplit[2])
		res += function(l, w, h)
	}

	return res
}

func findSurface(length int, width int, height int) int {
	lw := length * width
	wh := width * height
	hl := height * length

	return 2*lw + 2*wh + 2*hl + minOf(lw, wh, hl)
}

func findRibbon(length int, width int, height int) int {
	minVal := minOf(length, width, height)
	var secondMinVal int
	if minVal == length {
		secondMinVal = minOf(width, height)
	} else if minVal == width {
		secondMinVal = minOf(length, height)
	} else {
		secondMinVal = minOf(width, length)
	}

	wrapRibbon := minVal*2 + secondMinVal*2
	return length*width*height + wrapRibbon
}

func minOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

func readLines(path string) ([]string, error) {
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
