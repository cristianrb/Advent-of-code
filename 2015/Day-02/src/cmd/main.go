package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := readLines("inputs/input_01")
	res := findSurfaceFromList(data)
	print(res)
}

func findSurfaceFromList(data []string) int {
	res := 0
	for _, line := range data {
		lineSplit := strings.Split(line, "x")
		l, _ := strconv.Atoi(lineSplit[0])
		w, _ := strconv.Atoi(lineSplit[1])
		h, _ := strconv.Atoi(lineSplit[2])
		res += findSurface(l, w, h)
	}

	return res
}

func findSurface(length int, width int, height int) int {
	lw := length * width
	wh := width * height
	hl := height * length

	return 2*lw + 2*wh + 2*hl + minOf(lw, wh, hl)
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
