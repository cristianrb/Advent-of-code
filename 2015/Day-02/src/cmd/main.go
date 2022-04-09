package main

import (
	"cristianrbaoc2015d2/src/io"
	"cristianrbaoc2015d2/src/util"
	"strconv"
	"strings"
)

func main() {
	data, _ := io.ReadLines("inputs/input_01")
	res := findResFromList(data, findSurface)
	println(res)

	dataSecond, _ := io.ReadLines("inputs/input_02")
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

	return 2*lw + 2*wh + 2*hl + util.MinOf(lw, wh, hl)
}

func findRibbon(length int, width int, height int) int {
	minVal := util.MinOf(length, width, height)
	var secondMinVal int
	if minVal == length {
		secondMinVal = util.MinOf(width, height)
	} else if minVal == width {
		secondMinVal = util.MinOf(length, height)
	} else {
		secondMinVal = util.MinOf(width, length)
	}

	wrapRibbon := minVal*2 + secondMinVal*2
	return length*width*height + wrapRibbon
}
