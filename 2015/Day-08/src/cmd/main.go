package main

import (
	"cristianrb2015/io"
	"strconv"
)

func main() {
	lines, _ := io.ReadLines("inputs/input_01")
	total1 := countTotalCharactersDecoded(lines)
	println(total1)
	total2 := countTotalCharactersEncoded(lines)
	println(total2)

}

func countTotalCharactersDecoded(lines []string) int {
	total := 0
	for _, line := range lines {
		total += len(line) - len(unquote(line))
	}
	return total
}

func countTotalCharactersEncoded(lines []string) int {
	total := 0
	for _, line := range lines {
		total += len(quote(line)) - len(line)
	}
	return total
}

func unquote(str string) string {
	s, _ := strconv.Unquote(str)
	return s
}

func quote(str string) string {
	return strconv.Quote(str)
}
